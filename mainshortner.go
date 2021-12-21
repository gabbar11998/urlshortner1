package main

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/fs111/simpleconfig"
	godis "github.com/simonz05/godis/redis"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	COUNTER = "__counter__"
	HTTP    = "http"
)

var (
	redis        *godis.Client
