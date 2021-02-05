package main

import (
	"fmt"

	"github.com/sapiderman/seed-go/internal"
	log "github.com/sirupsen/logrus"
)

var (
	spashScreen = `                                         
	::::::::   :::::::::: :::::::::  :::     ::: :::::::::: :::::::::       :::    ::: :::::::::                                                    
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

	// start server
	internal.StartServer()
}
