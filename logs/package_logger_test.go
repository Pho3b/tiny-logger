package logs

//func TestPackageLoggerLogLvl(t *testing.T) {
//	assert.Equal(t, log_level.DebugLvl, packageLoggerConfigs.LogLvl.Lvl)
//	assert.Equal(t, log_level.DebugLvlName, packageLoggerConfigs.LogLvl.LvlName())
//
//	testLogsLvlVar1 := "MY_INSTANCE_LOGS_LVL"
//	testLogsLvlVar2 := "MY_INSTANCE_LOGS_LVL_2"
//
//	_ = os.Setenv(testLogsLvlVar1, string(log_level.WarnLvlName))
//	SetLogLvlEnvVariable(testLogsLvlVar1)
//	assert.Equal(t, log_level.WarnLvl, GetLogLvlIntValue())
//	assert.Equal(t, log_level.WarnLvlName, GetLogLvlName())
//
//	_ = os.Setenv(testLogsLvlVar2, string(log_level.InfoLvlName))
//	SetLogLvlEnvVariable(testLogsLvlVar2)
//	assert.Equal(t, log_level.InfoLvl, GetLogLvlIntValue())
//	assert.NotEqual(t, log_level.WarnLvl, GetLogLvlIntValue())
//
//	SetLogLvl(log_level.DebugLvlName)
//	_ = os.Unsetenv(testLogsLvlVar1)
//	_ = os.Unsetenv(testLogsLvlVar2)
//}
//
//func TestPackageLoggerGetLogLvlIntValue(t *testing.T) {
//	assert.Equal(t, log_level.DebugLvl, GetLogLvlIntValue())
//
//	SetLogLvl(log_level.ErrorLvlName)
//	assert.Equal(t, log_level.ErrorLvl, GetLogLvlIntValue())
//
//	testLogsLvlVar1 := "MY_INSTANCE_LOGS_LVL"
//	_ = os.Setenv(testLogsLvlVar1, string(log_level.InfoLvlName))
//	SetLogLvlEnvVariable(testLogsLvlVar1)
//	assert.Equal(t, log_level.InfoLvl, GetLogLvlIntValue())
//
//	SetLogLvl(log_level.DebugLvlName)
//	_ = os.Unsetenv(testLogsLvlVar1)
//}
//
//func TestPackageLogger_Info(t *testing.T) {
//	var buf bytes.Buffer
//	testLog := "my testing DEBUG log"
//	originalStdOut := os.Stdout
//	r, w, _ := os.Pipe()
//
//	os.Stdout = w
//	Info(testLog)
//
//	_ = w.Close()
//	_, _ = io.Copy(&buf, r)
//	os.Stdout = originalStdOut
//	assert.Contains(t, buf.String(), testLog)
//}
//
//func TestPackageLogger_InfoNotLogging(t *testing.T) {
//	var buf bytes.Buffer
//	testLog := "my testing DEBUG log"
//	originalStdOut := os.Stdout
//	r, w, _ := os.Pipe()
//
//	os.Stdout = w
//	logger := NewLogger()
//	logger.SetLogLvl(log_level.ErrorLvlName)
//	logger.Info(testLog)
//	logger.Warn(testLog)
//	logger.Debug(testLog)
//
//	_ = w.Close()
//	_, _ = io.Copy(&buf, r)
//	os.Stdout = originalStdOut
//	assert.NotContainsf(t, buf.String(), testLog, "logError-msg")
//}
//
//func TestPackageLogger_Debug(t *testing.T) {
//	var buf bytes.Buffer
//	testLog := "my testing INFO log"
//	originalStdOut := os.Stdout
//	r, w, _ := os.Pipe()
//
//	os.Stdout = w
//	Debug(testLog)
//
//	_ = w.Close()
//	_, _ = io.Copy(&buf, r)
//	os.Stdout = originalStdOut
//	assert.Contains(t, buf.String(), testLog)
//}
//
//func TestPackageLogger_Warn(t *testing.T) {
//	var buf bytes.Buffer
//	testLog := "my testing WARN log"
//	originalStdOut := os.Stdout
//	r, w, _ := os.Pipe()
//
//	os.Stdout = w
//	Warn(testLog)
//
//	_ = w.Close()
//	_, _ = io.Copy(&buf, r)
//	os.Stdout = originalStdOut
//	assert.Contains(t, buf.String(), testLog)
//}
//
//func TestPackageLogger_Error(t *testing.T) {
//	var buf bytes.Buffer
//	testLog := "my testing ERROR log"
//	originalStdErr := os.Stderr
//	r, w, _ := os.Pipe()
//
//	os.Stderr = w
//	Error(testLog)
//
//	_ = w.Close()
//	_, _ = io.Copy(&buf, r)
//	os.Stderr = originalStdErr
//	assert.Contains(t, buf.String(), testLog)
//}
//
//func TestPackageLogger_Log(t *testing.T) {
//	var buf bytes.Buffer
//	testLog := "my GENERIC log test"
//	originalStdOut := os.Stdout
//	r, w, _ := os.Pipe()
//
//	os.Stdout = w
//	Log("jldfald", testLog)
//	Log(colors.Black, testLog)
//	Log(colors.Cyan, testLog)
//
//	_ = w.Close()
//	_, _ = io.Copy(&buf, r)
//	os.Stdout = originalStdOut
//	assert.Contains(t, buf.String(), colors.White)
//	assert.Contains(t, buf.String(), colors.Black)
//	assert.Contains(t, buf.String(), colors.Cyan)
//}
