// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"github.com/google/logger"
	"github.com/spf13/viper"
	"os"
	"time"
)

// Info log
func Info(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BeaverBasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Beaver", false, false, lf)
		defer out.Close()

		out.Info(v...)
	}
}

// Infoln log
func Infoln(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BeaverBasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Beaver", false, false, lf)
		defer out.Close()

		out.Infoln(v...)
	}
}

// Infof log
func Infof(format string, v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BeaverBasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Beaver", false, false, lf)
		defer out.Close()

		out.Infof(format, v...)
	}
}

// Warning log
func Warning(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BeaverBasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Beaver", false, false, lf)
		defer out.Close()

		out.Warning(v...)
	}
}

// Warningln log
func Warningln(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BeaverBasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Beaver", false, false, lf)
		defer out.Close()

		out.Warningln(v...)
	}
}

// Warningf log
func Warningf(format string, v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BeaverBasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Beaver", false, false, lf)
		defer out.Close()

		out.Warningf(format, v...)
	}
}

// Error log
func Error(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning" || logLevel == "error"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BeaverBasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Beaver", false, false, lf)
		defer out.Close()

		out.Error(v...)
	}
}

// Errorln log
func Errorln(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning" || logLevel == "error"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BeaverBasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Beaver", false, false, lf)
		defer out.Close()

		out.Errorln(v...)
	}
}

// Errorf log
func Errorf(format string, v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning" || logLevel == "error"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BeaverBasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Beaver", false, false, lf)
		defer out.Close()

		out.Errorf(format, v...)
	}
}

// Fatal log
func Fatal(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning" || logLevel == "error" || logLevel == "fatal"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BeaverBasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Beaver", false, false, lf)
		defer out.Close()

		out.Fatal(v...)
	}
}

// Fatalln log
func Fatalln(v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning" || logLevel == "error" || logLevel == "fatal"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BeaverBasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Beaver", false, false, lf)
		defer out.Close()

		out.Fatalln(v...)
	}
}

// Fatalf log
func Fatalf(format string, v ...interface{}) {

	logLevel := viper.GetString("log.level")
	ok := logLevel == "info" || logLevel == "warning" || logLevel == "error" || logLevel == "fatal"

	if ok {
		currentTime := time.Now().Local()
		file := fmt.Sprintf(
			"%s%s/%s.log",
			os.Getenv("BeaverBasePath"),
			viper.GetString("log.path"),
			currentTime.Format("2006-01-02"),
		)
		lf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)

		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}

		defer lf.Close()

		out := logger.Init("Beaver", false, false, lf)
		defer out.Close()

		out.Fatalf(format, v...)
	}
}
