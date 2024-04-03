// Sample speech-quickstart uses the Google Cloud Speech API to transcribe
// audio.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
)

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// The path to the remote audio file to transcribe.
	// fileURI := "gs://cloud-samples-data/speech/brooklyn_bridge.raw"
	fn:="/home/siuyin/Music/audacity/export/turn_on_lights.mp3"
	// fn:="/home/siuyin/Music/audacity/export/increase_aircon.mp3"

	// Detects speech in the audio file.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_MP3,
			SampleRateHertz: 48000,
			LanguageCode:    "en-SG",
		},
		Audio: &speechpb.RecognitionAudio{
			// AudioSource: &speechpb.RecognitionAudio_Uri{Uri: fileURI},
			AudioSource: &speechpb.RecognitionAudio_Content{Content: audContent(fn)},
		},
	})
	if err != nil {
		log.Fatalf("failed to recognize: %v", err)
	}

	// Prints the results.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
}

func audContent(filename string) []byte {
	aud, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("audContent: %v", err)
	}
	return aud
}
