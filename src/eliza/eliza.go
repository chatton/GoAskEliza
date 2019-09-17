package eliza

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Eliza struct {
	// Eliza struct can be created using 2
	// interfaces, this allows you to provide multiple
	// different implementations of how you generate answers
	// and how you pick them.
	generator AnswerGenerator
	picker    AnswerPicker
	client  *mongo.Client
}

// NewEliza creates a new Eliza instance with teh given answer generator
// and answer picker.
func NewEliza(generator AnswerGenerator, picker AnswerPicker, client *mongo.Client) *Eliza {
	eliza := Eliza{generator: generator, picker: picker}
	eliza.client = client
	return &eliza
}

func (e *Eliza) saveQuestion(question string) error {
	ctx, _ := context.WithTimeout(context.TODO(), 2*time.Second)
	col := e.client.Database("eliza").Collection("questions")
	_, err := col.InsertOne(ctx, bson.M{
		"q": question,
	})
	if err != nil {
		return err
	}
	return nil
}

func (e *Eliza) saveAnswer(answer string) error {
	ctx, _ := context.WithTimeout(context.TODO(), 2*time.Second)
	col := e.client.Database("eliza").Collection("answers")
	_, err := col.InsertOne(ctx, bson.M{
		"a": answer,
	})
	if err != nil {
		return err
	}
	return nil
}

// GoAsk is the "main" exported function Eliza is needed for.
// it takes a single question in string format, and gives back
// a single response also in string format.
func (e *Eliza) GoAsk(question string) (string, error) {
	if err := e.saveQuestion(question); err != nil {
		return "", err
	}
	answers := e.generator.GenerateAnswers(question)
	answer := e.picker.PickAnswer(answers)

	if err := e.saveAnswer(answer); err != nil {
		return "", err
	}

	return answer, nil
}

// Questions returns a list of all asked questions
func (e *Eliza) Questions() ([]string, error) {
	ctx, _ := context.WithTimeout(context.TODO(), 2*time.Second)
	col := e.client.Database("eliza").Collection("questions")

	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	questions := make([]string, 0)
	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		questions = append(questions, result["q"].(string))
	}
	return questions, nil
}

// Answers returns a list of all given answers.
func (e *Eliza) Answers() ([]string, error) {
	ctx, _ := context.WithTimeout(context.TODO(), 2*time.Second)
	col := e.client.Database("eliza").Collection("answers")

	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	answers := make([]string, 0)
	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		answers = append(answers, result["a"].(string))
	}
	return answers, nil
}

// https://github.com/golang/go/wiki/CodeReviewComments#interfaces
// The documentation states that interfaces belong in the package that is
// going to use the interface type, not in with the implementations.

// AnswerGenerator should be able to give back a slice of answers when given a question.
type AnswerGenerator interface {
	GenerateAnswers(question string) []string
}

// AnswerPicker is in charge of picking a response out of a list of answers.
type AnswerPicker interface {
	PickAnswer(answers []string) string
}
