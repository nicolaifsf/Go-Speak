# Go-Speak
Speech to Text for Golang. Capable of continuous speech. Also wrappers for various Speech to Text APIs

package nicolaifsf/go-speak imported as speech implements continuous speech recognition and offers wrappers for Google Speech, Wit.ai, IBM Speech to text, and AT&T Speech to Text apis.

## Usage

### Basic Usage
```go
import("github.com/nicolaifsf/go-speak")
//example
speech.GetWitKey() //note that this package is imported as speech
```

### Wit.AI
###### Converting a sound file to text using wit.ai
```go
func main(){
  speech.SetWitKey("/**Your Wit API Key here**/") //Wit API Key MUST be set before calling any other Wit.AI functions
  speech.SendWitVoice("test.wav")
}
```

### Continuous Speech Recognition
```go
func main(){
  speech.SetWitKey("/**Your Wit API Key here**/") //Currently ContinuousRecognition() uses wit.ai for speech recognition
  speech.ContinuousRecognition()
}
```
ContinuousRecognition() currently simply prints out the json response from wit.ai call
## Installation
