package main

import (
	"context"
	"fmt"

	"github.com/sapiderman/seed-go/internal"
	"github.com/sapiderman/seed-go/internal/config"
	log "github.com/sirupsen/logrus"
)

var (
	spashScreen = `                                         
	::::::::  :::::::::: :::::::::  :::     ::: :::::::::: :::::::::        :::    ::: :::::::::                                                    
	:+:    :+: :+:        :+:    :+: :+:     :+: :+:        :+:    :+:      :+:    :+: :+:    :+:                                                   
	+:+        +:+        +:+    +:+ +:+     +:+ +:+        +:+    +:+      +:+    +:+ +:+    +:+
	+#++:++#++ +#++:++#   +#++:++#:  +#+     +:+ +#++:++#   +#++:++#:       +#+    +:+ +#++:++#+ 
	       +#+ +#+        +#+    +#+  +#+   +#+  +#+        +#+    +#+      +#+    +#+ +#+ 
	#+#    #+# #+#        #+#    #+#   #+#+#+#   #+#        #+#    #+#      #+#    #+# #+# 
	 ########  ########## ###    ###     ###     ########## ###    ###       ########  ### 	                                                   							 
	`
)

func init() {
	fmt.Println(spashScreen)
	log.Info("intializing")
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

// Main entry point
func main() {

	// create config and config context
	cfg := config.LoadConfig()
	cfgCtxKey := internal.ContextKey(internal.ConfigKey)
	cfgContext := context.WithValue(context.Background(), cfgCtxKey, cfg)

	// start server
	server := internal.NewServer(cfgContext)
	server.StartServer(cfgContext)
}
