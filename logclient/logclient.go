package logclient

// Output channel to LogClient
var LogOutputChan = make(chan []Entry, DEFAULT_INPUT_CHANNEL_BUF_CAPACITY)

type LogClient struct {
	conf       *Configuration
	logReader  *LogReader
	networkOut *NetworkOut
	// LogClient input channel to which an application sends input entries.
	InputChan chan []Entry
	// Log events pass through LogClient when their log level is > minCollectLogLevel.
	minCollectLogLevel int
}

func NewLogClient(conf *Configuration, passedInputChan chan []Entry) (*LogClient, error) {
	// initialize operators top down
	var inputChan chan []Entry
	ntkOut, err := NewNetworkOut(conf)
	if err != nil {
		return nil, err
	}
	inputChan = ntkOut.inputChan

	// bottom most operator
	// If the given passedInputChan is not null, pass to
	// the bottom operator.
	var logR *LogReader
	logR, err = NewLogReader(conf, ntkOut.inputChan, passedInputChan)
	if err != nil {
		return nil, err
	}
	inputChan = logR.inputChan

	// Return the execution object instance.
	return &LogClient{
		minCollectLogLevel: MininumCollectionLogLevel,
		conf:               conf,
		logReader:          logR,
		networkOut:         ntkOut,
		InputChan:          inputChan,
	}, nil
}

func (lc *LogClient) Close() {
	lc.logReader.Close()
	lc.networkOut.Close()
}

func (lc *LogClient) LogReader() *LogReader {
	return lc.logReader
}

func (lc *LogClient) NetworkOut() *NetworkOut {
	return lc.networkOut
}
