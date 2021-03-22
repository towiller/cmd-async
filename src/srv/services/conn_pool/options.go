package connpool

import (
	"time"
)

type HttpOptions struct {
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     int
	DialTimeout         int
	KeepAlive           int
}

type HttpOption func(o *HttpOptions)

type MysqlOptions struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DefaultDb    string `json:"default_db"`
	Charset      string `json:"charset"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

type MysqlOption func(o *MysqlOptions)

func MysqlHost(host string) MysqlOption {
	return func(o *MysqlOptions) {
		o.Host = host
	}
}

func MysqlPort(port int) MysqlOption {
	return func(o *MysqlOptions) {
		o.Port = port
	}
}

func MysqlUserName(username string) MysqlOption {
	return func(o *MysqlOptions) {
		o.Username = username
	}
}

func MysqlPassword(password string) MysqlOption {
	return func(o *MysqlOptions) {
		o.Password = password
	}
}

func MysqlDefaultDb(defaultDb string) MysqlOption {
	return func(o *MysqlOptions) {
		o.DefaultDb = defaultDb
	}
}

func MysqlCharset(charset string) MysqlOption {
	return func(o *MysqlOptions) {
		o.Charset = charset
	}
}

func MysqlMaxOpenConns(maxOpenConns int) MysqlOption {
	return func(o *MysqlOptions) {
		o.MaxOpenConns = maxOpenConns
	}
}

func MysqlMaxIdleConns(maxIdleConns int) MysqlOption {
	return func(o *MysqlOptions) {
		o.MaxIdleConns = maxIdleConns
	}
}

type RedisOptions struct {
	Host        string        `json:"host"`
	Port        int           `json:"port"`
	Password    string        `json:"password"`
	MaxIdle     int           `json:"max_idle"`
	IdleTimeout time.Duration `json:"idle_timeout"`
	SelectDB    int           `json:"select_db"`
	DbCount     int           `json:"db_count"`
}

type RedisOption func(*RedisOptions)

func RedisHost(host string) RedisOption {
	return func(o *RedisOptions) {
		o.Host = host
	}
}

func RedisPort(port int) RedisOption {
	return func(o *RedisOptions) {
		o.Port = port
	}
}

func RedisPassword(password string) RedisOption {
	return func(o *RedisOptions) {
		o.Password = password
	}
}

func RedisMaxIdle(maxidle int) RedisOption {
	return func(o *RedisOptions) {
		o.MaxIdle = maxidle
	}
}

func RedisIdleTimeout(idleTimeout time.Duration) RedisOption {
	return func(o *RedisOptions) {
		o.IdleTimeout = idleTimeout * time.Second
	}
}

func RedisSelectDB(selectDb int) RedisOption {
	return func(o *RedisOptions) {
		o.SelectDB = selectDb
	}
}

func RedisDbCount(dbCount int) RedisOption {
	return func(o *RedisOptions) {
		o.DbCount = dbCount
	}
}
