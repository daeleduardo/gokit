package lib

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)
//Most Common Exit Status Code of bash command
type CmdExitCode struct {
	Success                                     int `default:"0"`
	OperationNotPermitted                       int `default:"1"`
	NoSuchFileOrDirectory                       int `default:"2"`
	NoSuchProcess                               int `default:"3"`
	InterruptedSystemCall                       int `default:"4"`
	InputOutputError                            int `default:"5"`
	NoSuchDeviceOrAddress                       int `default:"6"`
	ArgumentListTooLong                         int `default:"7"`
	ExecFormatError                             int `default:"8"`
	BadFileDescriptor                           int `default:"9"`
	NoChildProcesses                            int `default:"10"`
	ResourceTemporarilyUnavailable              int `default:"11"`
	CannotAllocateMemory                        int `default:"12"`
	PermissionDenied                            int `default:"13"`
	BadAddress                                  int `default:"14"`
	BlockDeviceRequired                         int `default:"15"`
	DeviceOrResourceBusy                        int `default:"16"`
	FileExists                                  int `default:"17"`
	InvalidCrossDeviceLink                      int `default:"18"`
	NoSuchDevice                                int `default:"19"`
	NotADirectory                               int `default:"20"`
	IsADirectory                                int `default:"21"`
	InvalidArgument                             int `default:"22"`
	TooManyOpenFilesInSystem                    int `default:"23"`
	TooManyOpenFiles                            int `default:"24"`
	InappropriateIoctlForDevice                 int `default:"25"`
	TextFileBusy                                int `default:"26"`
	FileTooLarge                                int `default:"27"`
	NoSpaceLeftOnDevice                         int `default:"28"`
	IllegalSeek                                 int `default:"29"`
	ReadOnlyFileSystem                          int `default:"30"`
	TooManyLinks                                int `default:"31"`
	BrokenPipe                                  int `default:"32"`
	NumericalArgumentOutOfDomain                int `default:"33"`
	NumericalResultOutOfRange                   int `default:"34"`
	ResourceDeadlockAvoided                     int `default:"35"`
	FileNameTooLong                             int `default:"36"`
	NoLocksAvailable                            int `default:"37"`
	FunctionNotImplemented                      int `default:"38"`
	DirectoryNotEmpty                           int `default:"39"`
	TooManyLevelsOfSymbolicLinks                int `default:"40"`
	NoMessageOfDesiredType                      int `default:"42"`
	IdentifierRemoved                           int `default:"43"`
	ChannelNumberOutOfRange                     int `default:"44"`
	Level2NotSynchronized                       int `default:"45"`
	Level3Halted                                int `default:"46"`
	Level3Reset                                 int `default:"47"`
	LinkNumberOutOfRange                        int `default:"48"`
	ProtocolDriverNotAttached                   int `default:"49"`
	NoCsiStructureAvailable                     int `default:"50"`
	Level2Halted                                int `default:"51"`
	InvalidExchange                             int `default:"52"`
	InvalidRequestDescriptor                    int `default:"53"`
	ExchangeFull                                int `default:"54"`
	NoAnode                                     int `default:"55"`
	InvalidRequestCode                          int `default:"56"`
	InvalidSlot                                 int `default:"57"`
	BadFontFileFormat                           int `default:"59"`
	DeviceNotAStream                            int `default:"60"`
	NoDataAvailable                             int `default:"61"`
	TimerExpired                                int `default:"62"`
	OutOfStreamsResources                       int `default:"63"`
	MachineIsNotOnTheNetwork                    int `default:"64"`
	PackageNotInstalled                         int `default:"65"`
	ObjectIsRemote                              int `default:"66"`
	LinkHasBeenSevered                          int `default:"67"`
	AdvertiseError                              int `default:"68"`
	SrmountError                                int `default:"69"`
	CommunicationErrorOnSend                    int `default:"70"`
	ProtocolError                               int `default:"71"`
	MultihopAttempted                           int `default:"72"`
	RfsSpecificError                            int `default:"73"`
	BadMessage                                  int `default:"74"`
	ValueTooLargeForDefinedDataType             int `default:"75"`
	NameNotUniqueOnNetwork                      int `default:"76"`
	FileDescriptorInBadState                    int `default:"77"`
	RemoteAddressChanged                        int `default:"78"`
	CanNotAccessANeededSharedLibrary            int `default:"79"`
	AccessingACorruptedSharedLibrary            int `default:"80"`
	LibSectionInAOutCorrupted                   int `default:"81"`
	AttemptingToLinkInTooManySharedLibraries    int `default:"82"`
	CannotExecASharedLibraryDirectly            int `default:"83"`
	InvalidOrIncompleteMultibyteOrWideCharacter int `default:"84"`
	InterruptedSystemCallShouldBeRestarted      int `default:"85"`
	StreamsPipeError                            int `default:"86"`
	TooManyUsers                                int `default:"87"`
	SocketOperationOnNonSocket                  int `default:"88"`
	DestinationAddressRequired                  int `default:"89"`
	MessageTooLong                              int `default:"90"`
	ProtocolWrongTypeForSocket                  int `default:"91"`
	ProtocolNotAvailable                        int `default:"92"`
	ProtocolNotSupported                        int `default:"93"`
	SocketTypeNotSupported                      int `default:"94"`
	OperationNotSupported                       int `default:"95"`
	ProtocolFamilyNotSupported                  int `default:"96"`
	AddressFamilyNotSupportedByProtocol         int `default:"97"`
	AddressAlreadyInUse                         int `default:"98"`
	CannotAssignRequestedAddress                int `default:"99"`
	NetworkIsDown                               int `default:"100"`
	NetworkIsUnreachable                        int `default:"101"`
	NetworkDroppedConnectionOnReset             int `default:"102"`
	SoftwareCausedConnectionAbort               int `default:"103"`
	ConnectionResetByPeer                       int `default:"104"`
	NoBufferSpaceAvailable                      int `default:"105"`
	TransportEndpointIsAlreadyConnected         int `default:"106"`
	TransportEndpointIsNotConnected             int `default:"107"`
	CannotSendAfterTransportEndpointShutdown    int `default:"108"`
	TooManyReferences                           int `default:"109"`
	ConnectionTimedOut                          int `default:"110"`
	ConnectionRefused                           int `default:"111"`
	HostIsDown                                  int `default:"112"`
	NoRouteToHost                               int `default:"113"`
	OperationAlreadyInProgress                  int `default:"114"`
	OperationNowInProgress                      int `default:"115"`
	StaleFileHandle                             int `default:"116"`
	StructureNeedsCleaning                      int `default:"117"`
	NotAXenixNamedTypeFile                      int `default:"118"`
	NoXenixSemaphoresAvailable                  int `default:"119"`
	IsANamedTypeFile                            int `default:"120"`
	RemoteIOError                               int `default:"121"`
	DiskQuotaExceeded                           int `default:"122"`
	NoMediumFound                               int `default:"123"`
	OperationCanceled                           int `default:"125"`
	RequiredKeyNotAvailable                     int `default:"126"`
	KeyHasExpired                               int `default:"127"`
	KeyHasBeenRevoked                           int `default:"128"`
	KeyWasRejectedByService                     int `default:"129"`
	OwnerDied                                   int `default:"130"`
	StateNotRecoverable                         int `default:"131"`
	OperationNotPossibleDueToRfKill             int `default:"132"`
	MemoryPageHasHardwareError                  int `default:"133"`
}

type CommandResult struct {
	Stdout string
	Stderr string
	ExitCode int
}

//Wrapper to execute bash commands and scripts
func ShellCommand(command string) (CommandResult, error) {
	
	result := CommandResult{}
	cmd := exec.Command("bash", "-c", command)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()


	if err != nil {
		result.Stdout = ""
		result.Stderr = strings.TrimSpace(stderr.String())
		result.ExitCode , _ = strconv.Atoi(err.Error())
		return  result , err
	}

	result.Stdout = strings.TrimSpace(stdout.String())
	return  result , nil
}