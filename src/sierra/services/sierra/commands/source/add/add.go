package add

import "fmt"

type Args struct {
	Positional struct {
		Source string `positional-arg-name:"<source>" required:"true"`
	} `positional-args:"yes" required:"yes"`
}

func Run(args Args) error {
	fmt.Println("add", args.Positional.Source)
	return nil
}
