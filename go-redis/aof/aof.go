package aof

import (
	"go-redis/config"
	"go-redis/database"
	"go-redis/lib/logger"
	"go-redis/lib/utils"
	"go-redis/resp/reply"
	"os"
	"strconv"
)

type CmdLine = [][]byte

const aofBufferSize = 1 << 16

type payload struct {
	cmdLine CmdLine
	dbIndex int
}

type AofHandler struct {
	database    database.Database
	aofChan     chan *payload
	aofFile     *os.File
	aofFilename string
	currentDB   int
}

// NewAofHandler
func NewAofHandler(database database.Database) (*AofHandler, error) {
	handler := &AofHandler{}
	handler.aofFilename = config.Properties.AppendFilename
	handler.database = database
	handler.LoadAof()
	aofile, err := os.OpenFile(handler.aofFilename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	handler.aofFile = aofile
	handler.aofChan = make(chan *payload, aofBufferSize)
	go func() {
		handler.handleAof()
	}()
	return handler, nil
}

// Add payload(set k v) -> aofChan
func (handler *AofHandler) AddAof(dbIndex int, cmd CmdLine) {
	if config.Properties.AppendOnly && handler.aofChan != nil {
		handler.aofChan <- &payload{
			cmdLine: cmd,
			dbIndex: dbIndex,
		}
	}

}

// handleAof payload(set k v) <- aofChan(落盘)
func (handler *AofHandler) handleAof() {
	handler.currentDB = 0
	for p := range handler.aofChan {
		if p.dbIndex != handler.currentDB {
			data := reply.MakeMultiBulkReply(utils.ToCmdLine("select", strconv.Itoa(p.dbIndex))).ToBytes()
			_, err := handler.aofFile.Write(data)
			if err != nil {
				logger.Error(err)
				continue
			}
			handler.currentDB = p.dbIndex
		}
		data := reply.MakeMultiBulkReply(p.cmdLine).ToBytes()
		_, err := handler.aofFile.Write(data)
		if err != nil {
			logger.Error(err)
		}
	}
}

// LoadAof
func (handler *AofHandler) LoadAof() {

}
