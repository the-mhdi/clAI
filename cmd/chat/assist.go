package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/the-mhdi/maShit/util"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func loadModel(conf *util.Config) error {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(conf.GEMINI_API_KEY))

	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	defer client.Close()

	model := client.GenerativeModel(conf.Model)
	model.SetTemperature(conf.Temperature)
	model.SetTopK(conf.TopK)
	model.SetTopP(conf.TopP)
	model.SetMaxOutputTokens(conf.MaxOutput)
	model.ResponseMIMEType = conf.ResponseMIMEType

	session := model.StartChat()
	session.History = []*genai.Content{
		{
			Role: "user",
			Parts: []genai.Part{
				genai.Text("you are an expert in coding and network engineering, you have great and vast knowledge of decentralized networks like ethereum and bitcoin and years of experience coding in C, Go lang and Solidity smart contract.\n\nyou are incredibly capable of explaining complex programming and decentralized networking concepts in simple language. \nyou've spent years developing decentralized applications (Dapps) on Ethereum and been contributing to the Go-ethereum project over the years and now you can completely analyze the source code of this project and explain the logic behind every line of code written there.\n\nyou have extreme knowledge of crypto-economic, tokecnomics. \n\nYou pride yourself on incredible accuracy and attention to detail. you are a wonderful and patient tutor who explains and teaches coding problems thoroughly and in great detail and you're also an amazing programmer and coder.\n\nyour job is to listen to users' questions and problems and provide detailed answers and code snippets if possible and also help user learn those concepts along the way."),
			},
		},
		{
			Role: "model",
			Parts: []genai.Part{
				genai.Text("Understood.  I'm ready to put on my expert coding and network engineering hat and assist you with your questions about decentralized networks, programming (especially in C, Go, and Solidity), and crypto-economics.  Ask away! I will do my best to break down complex topics into easily digestible explanations, provide code examples when relevant, and guide you through the learning process. I'll leverage my \"knowledge\" (as trained on a massive dataset) and Python tools to assist with precise computations and analyses. Remember, I'm here to help you understand, so don't hesitate to ask clarifying questions if anything is unclear.\n"),
			},
		},
	}
	run(model, ctx)

	return errors.New("some issue running the model")
}

func run(model *genai.GenerativeModel, ctx context.Context) {

	for {
		input, err := util.GetStdIn()

		if err != nil {
			fmt.Print(err)
		}

		iter := model.GenerateContentStream(ctx, genai.Text(input))

		for {
			resp, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			util.PrintResponse(resp)
		}
	}
}
