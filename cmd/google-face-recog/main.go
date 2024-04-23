// Uses the Google Cloud Vision API to detect faces
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vision "cloud.google.com/go/vision/v2/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
)

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name of the image file to annotate.
	fn1 := "/home/siuyin/Downloads/old/face-20240423-120243.jpg"

	req := &visionpb.BatchAnnotateImagesRequest{
		Requests: []*visionpb.AnnotateImageRequest{
			{
				Image:    &visionpb.Image{Content: imgContent(fn1)},
				Features: []*visionpb.Feature{{Type: visionpb.Feature_FACE_DETECTION}},
			},
		},
	}

	resp, err := client.BatchAnnotateImages(ctx, req)
	if err != nil {
		log.Fatalf("BatchAnnotateImages: %v", err)
	}

	for _, r := range resp.Responses {
		// fmt.Printf("----\n%s\n",r.FullTextAnnotation.Text)
		fmt.Printf("----\n%s\n", r.GetFaceAnnotations())
	}
}

func imgContent(filename string) []byte {
	imgBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	return imgBytes
}
