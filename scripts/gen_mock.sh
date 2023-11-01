#!/bin/sh

go generate ./...

uName=`uname -s`
osName=${uName: 0: 4}

for path in $(find .  -name "mock_*.go")
do
  echo $path

  if [ "${osName}" == "Darw" ] # Darwin (mac os)
  then
    sed -i "" '/mock.Mock.Test(t)/d' $path
    sed -i "" '/t.Cleanup(func() { mock.AssertExpectations(t) })/d' $path
  elif  [ "${osName}" == "Linu" ] # Linux
  then
    sed -i "" '/mock.Mock.Test(t)/d' $path
    sed -i "" '/t.Cleanup(func() { mock.AssertExpectations(t) })/d' $path
  elif [ "${osName}" == "MING" ] # windows, git-bash
  then
    sed -i '/mock.Mock.Test(t)/d' $path
    sed -i '/t.Cleanup(func() { mock.AssertExpectations(t) })/d' $path
  fi

done