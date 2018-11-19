package writer


type StubWriter struct{
	arrayOfOutputs []string
}

func (stubWriter *StubWriter) Write(text string) {
	stubWriter.arrayOfOutputs = append(stubWriter.arrayOfOutputs, text)
}

func (stubWriter *StubWriter) WriteWarning(text string) {
	stubWriter.arrayOfOutputs = append(stubWriter.arrayOfOutputs, text)
}

func (stubWriter *StubWriter) CleanOutputs() {
	stubWriter.arrayOfOutputs = []string{}
}

func (stubWriter *StubWriter) GetOutputs() []string {
	return stubWriter.arrayOfOutputs
}

func (stubWriter StubWriter) GetLastMessage() string {
	return stubWriter.GetOutputs()[len(stubWriter.arrayOfOutputs)-1]
}