package main

import (
	"context"
	"fmt"

	"github.com/sapiderman/test-seed/internal"
	"github.com/sapiderman/test-seed/internal/config"
	log "github.com/sirupsen/logrus"
)

var (
	spashScreen = `                                         
	::::::::  :::::::::: :::::::::  :::     ::: :::::::::: :::::::::       
	:+:    :+: :+:        :+:    :+: :+:     :+: :+:        :+:    :+:      
	+:+        +:+        +:+    +:+ +:+     +:+ +:+        +:+    +:+      
	+#++:++#++ +#++:++#   +#++:++#:  +#+     +:+ +#++:++#   +#++:++#:       
		   +#+ +#+        +#+    +#+  +#+   +#+  +#+        +#+    +#+      
	#+#    #+# #+#        #+#    #+#   #+#+#+#   #+#        #+#    #+#      
	 ########  ########## ###    ###     ###     ########## ###    ###      
	:::    ::: :::::::::                                                    
	:+:    :+: :+:    :+:                                                   
	+:+    +:+ +:+    +:+                                                   
	+#+    +:+ +#++:++#+                                                    
	+#+    +#+ +#+                                                          
	#+#    #+# #+#                                                          
	 ########  ###                                                          							 
	`
)

func init() {
	fmt.Println(spashScreen)
	log.Info("intializing")
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	return
}

//Main entry point
func main() {

	var cfg *config.Configuration

	// create config and app context
	cfg = config.LoadConfig()
	cfgCtxKey := internal.ContextKey(internal.ConfigKey)
	appContext := context.WithValue(context.Background(), cfgCtxKey, cfg)

	// start server
	server := internal.NewServer(appContext)
	server.StartServer(appContext)
}
