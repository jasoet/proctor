package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"sync"
	"time"
)

func GetStringDefault(viper *viper.Viper, key string, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	return viper.GetString(key)
}

func GetInt64Ref(viper *viper.Viper, key string) *int64 {
	value := viper.GetInt64(key)
	return &value
}

func GetInt32Ref(viper *viper.Viper, key string) *int32 {
	value := viper.GetInt32(key)
	return &value
}

func GetMapFromJson(viper *viper.Viper, key string) map[string]string {
	var jsonStr = []byte(viper.GetString(key))
	var annotations map[string]string

	err := json.Unmarshal(jsonStr, &annotations)
	if err != nil {
		_ = fmt.Errorf("invalid Value for key %s, errors %v", key, err.Error())
	}

	return annotations
}

var once sync.Once
var config ProctorConfig

type ProctorConfig struct {
	viper                            *viper.Viper
	KubeConfig                       string
	KubeContext                      string
	LogLevel                         string
	AppPort                          string
	DefaultNamespace                 string
	RedisAddress                     string
	LogsStreamReadBufferSize         int
	RedisMaxActiveConnections        int
	LogsStreamWriteBufferSize        int
	KubePodsListWaitTime             time.Duration
	KubeLogProcessWaitTime           time.Duration
	KubeJobActiveDeadlineSeconds     *int64
	KubeJobRetries                   *int32
	PostgresUser                     string
	PostgresPassword                 string
	PostgresHost                     string
	PostgresPort                     int
	AuthPluginExported               string
	PostgresDatabase                 string
	PostgresMaxConnections           int
	PostgresConnectionMaxLifetime    int
	NewRelicAppName                  string
	NewRelicLicenceKey               string
	MinClientVersion                 string
	ScheduledJobsFetchIntervalInMins int
	MailUsername                     string
	MailServerHost                   string
	MailPassword                     string
	MailServerPort                   string
	JobPodAnnotations                map[string]string
	SentryDSN                        string
	DocsPath                         string
	AuthPluginBinary                 string
}

func Load() ProctorConfig {
	fang := viper.New()
	fang.AutomaticEnv()
	fang.SetEnvPrefix("PROCTOR")
	proctorConfig := ProctorConfig{
		viper:                            fang,
		KubeConfig:                       fang.GetString("KUBE_CONFIG"),
		KubeContext:                      GetStringDefault(fang, "KUBE_CONTEXT", "default"),
		LogLevel:                         GetStringDefault(fang, "LOG_LEVEL", "DEBUG"),
		AppPort:                          GetStringDefault(fang, "APP_PORT", "5001"),
		DefaultNamespace:                 fang.GetString("DEFAULT_NAMESPACE"),
		RedisAddress:                     fang.GetString("REDIS_ADDRESS"),
		RedisMaxActiveConnections:        fang.GetInt("REDIS_MAX_ACTIVE_CONNECTIONS"),
		LogsStreamReadBufferSize:         fang.GetInt("LOGS_STREAM_READ_BUFFER_SIZE"),
		LogsStreamWriteBufferSize:        fang.GetInt("LOGS_STREAM_WRITE_BUFFER_SIZE"),
		KubePodsListWaitTime:             time.Duration(fang.GetInt("KUBE_POD_LIST_WAIT_TIME")),
		KubeLogProcessWaitTime:           time.Duration(fang.GetInt("KUBE_LOG_PROCESS_WAIT_TIME")),
		KubeJobActiveDeadlineSeconds:     GetInt64Ref(fang, "KUBE_JOB_ACTIVE_DEADLINE_SECONDS"),
		KubeJobRetries:                   GetInt32Ref(fang, "KUBE_JOB_RETRIES"),
		PostgresUser:                     fang.GetString("POSTGRES_USER"),
		PostgresPassword:                 fang.GetString("POSTGRES_PASSWORD"),
		PostgresHost:                     fang.GetString("POSTGRES_HOST"),
		PostgresPort:                     fang.GetInt("POSTGRES_PORT"),
		PostgresDatabase:                 fang.GetString("POSTGRES_DATABASE"),
		PostgresMaxConnections:           fang.GetInt("POSTGRES_MAX_CONNECTIONS"),
		PostgresConnectionMaxLifetime:    fang.GetInt("POSTGRES_CONNECTIONS_MAX_LIFETIME"),
		NewRelicAppName:                  fang.GetString("NEW_RELIC_APP_NAME"),
		NewRelicLicenceKey:               fang.GetString("NEW_RELIC_LICENCE_KEY"),
		MinClientVersion:                 fang.GetString("MIN_CLIENT_VERSION"),
		ScheduledJobsFetchIntervalInMins: fang.GetInt("SCHEDULED_JOBS_FETCH_INTERVAL_IN_MINS"),
		MailUsername:                     fang.GetString("MAIL_USERNAME"),
		MailServerHost:                   fang.GetString("MAIL_SERVER_HOST"),
		MailPassword:                     fang.GetString("MAIL_PASSWORD"),
		MailServerPort:                   fang.GetString("MAIL_SERVER_PORT"),
		JobPodAnnotations:                GetMapFromJson(fang, "JOB_POD_ANNOTATIONS"),
		SentryDSN:                        fang.GetString("SENTRY_DSN"),
		DocsPath:                         fang.GetString("DOCS_PATH"),
		AuthPluginBinary:                 fang.GetString("AUTH_PLUGIN_BINARY"),
		AuthPluginExported:               GetStringDefault(fang, "AUTH_PLUGIN_EXPORTED", "Auth"),
	}

	return proctorConfig
}

var reset = false

func Reset() {
	reset = true
}

func Config() ProctorConfig {
	once.Do(func() {
		config = Load()
	})

	if reset {
		config = Load()
		reset = false
	}

	return config
}
