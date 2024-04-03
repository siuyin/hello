// Sample vision-quickstart uses the Google Cloud Vision API to label an image.
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
	filename := "/home/siuyin/Downloads/old/increaseAircon IMG_20240403_112530~2.jpg"

	// file, err := os.Open(filename)
	// if err != nil {
	// 	log.Fatalf("Failed to read file: %v", err)
	// }
	// defer file.Close()

	imgBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	req := &visionpb.BatchAnnotateImagesRequest{
		Requests: []*visionpb.AnnotateImageRequest{{
			Image: &visionpb.Image{Content: imgBytes},
			Features: []*visionpb.Feature{{Type: visionpb.Feature_DOCUMENT_TEXT_DETECTION}},
		}},
	}

	resp, err := client.BatchAnnotateImages(ctx, req)
	if err != nil {
		log.Fatalf("BatchAnnotateImages: %v", err)
	}

	for _, r := range resp.Responses {
		fmt.Printf(r.FullTextAnnotation.Text)
	}
}
