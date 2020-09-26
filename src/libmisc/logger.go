package libmisc

import (
	"fmt"
	"log"
	"os"
)

var (
    warningLogger *log.Logger
    infoLogger    *log.Logger
    errorLogger   *log.Logger
)

func SetLoggerFilePath( logfilepath string ) {
    file, err := os.OpenFile( logfilepath  , os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Println(err)
        return
    }

    infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Log( level string , format string, msgs... interface{} ) {
    if infoLogger == nil || warningLogger == nil || errorLogger == nil {
        log.Println( "logger doesn't be initialized" )
        return
    }

    switch level {
    case "I":
        infoLogger.Println( fmt.Sprintf(format,msgs...) )
    case "W":
        warningLogger.Println( fmt.Sprintf(format,msgs...) )
    case "E":
        errorLogger.Println( fmt.Sprintf(format,msgs...) )
    default:
        errorLogger.Println( "log level only support:I|W|E" )
    }

}

