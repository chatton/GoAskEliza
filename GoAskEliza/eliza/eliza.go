package eliza

type Eliza struct {

}


func New() Eliza {
    eliza := Eliza{}
    return eliza
} 

func (*Eliza) GoAsk(question string) string {
    return "You asked: " + question
} 