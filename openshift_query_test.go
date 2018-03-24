package main

import (
	_ "testing"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompleteArgs(t *testing.T) {

	osArgs := [] string {"wait-for-build", "-appname=abc", "-username=user1", "-password=pass1"}
	queryArgs := constructArgs(osArgs)

	assert.Equal(t, "wait-for-build", queryArgs.queryName)
	require.Equal(t, "wait-for-build", queryArgs.queryName)

	assert.Equal(t, "user1", queryArgs.userName)
	assert.Equal(t, "pass1", queryArgs.password)
	assert.Equal(t, "abc", queryArgs.appName)
}

func TestNoArgs(t *testing.T) {
	osArgs := [] string {"wait-for-build"}
	queryArgs := constructArgs(osArgs)

	assert.Equal(t, "dev", queryArgs.userName)
	assert.Equal(t, "dev", queryArgs.password)
	assert.Equal(t, "pet-clinic", queryArgs.appName)
}

func TestGetBuildStatus(t *testing.T) {
	osArgs := [] string {"wait-for-build"}
	constructArgs(osArgs)
	status := getBuildStatus("pet-clinic")
	assert.Equal(t, "Complete", status)
}

func TestMain(t *testing.T) {

}