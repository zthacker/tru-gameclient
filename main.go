package main

import (
	"gameclient/client"
	"gameclient/frontend"
	"gameclient/proto"
	"github.com/andrew-d/go-termutil"
	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"regexp"
)

const (
	backgroundColor = tcell.Color234
	textColor       = tcell.ColorWhite
	fieldColor      = tcell.Color24
)

type connectInfo struct {
	PlayerName string
	Address    string
	Password   string
}

func loginApp(info *connectInfo) *tview.Application {
	app := tview.NewApplication()
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow)
	flex.SetBorder(true).
		SetTitle(" Connect to TR Ultimate server ").
		SetBackgroundColor(backgroundColor)
	errors := tview.NewTextView().
		SetText(" Use the tab key to change fields, and enter to submit")
	errors.SetBackgroundColor(backgroundColor)
	form := tview.NewForm()
	re := regexp.MustCompile("^[a-zA-Z0-9]+$")
	form.AddInputField("Player Name", "", 16, func(textCheck string, lastChar rune) bool {
		result := re.MatchString(textCheck)
		if !result {
			errors.SetText(" Only alphanumeric characters are allowed")
		}
		return result
	}, nil).
		AddInputField("Server Address", ":8888", 32, nil, nil).
		AddPasswordField("Server Password", "", 32, '*', nil).
		AddButton("Connect", func() {
			info.PlayerName = form.GetFormItem(0).(*tview.InputField).GetText()
			info.Address = form.GetFormItem(1).(*tview.InputField).GetText()
			info.Password = form.GetFormItem(2).(*tview.InputField).GetText()
			if info.PlayerName == "" || info.Address == "" {
				errors.SetText(" All fields are required.")
				return
			}
			app.Stop()
		}).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetLabelColor(textColor).
		SetButtonBackgroundColor(fieldColor).
		SetFieldBackgroundColor(fieldColor).
		SetBackgroundColor(backgroundColor)
	flex.AddItem(errors, 1, 1, false)
	flex.AddItem(form, 0, 1, false)
	app.SetRoot(flex, true).SetFocus(form)
	return app
}

func main() {

	if !termutil.Isatty(os.Stdin.Fd()) {
		logrus.Fatal("this program must be run in a terminal")
	}

	info := connectInfo{}
	app := loginApp(&info)
	app.Run()

	view := frontend.NewView()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(info.Address, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcClient := proto.NewGameBackendClient(conn)
	c := client.NewGameClient()
	playerID := uuid.New()
	logrus.Info(playerID)

	c.Connect(grpcClient, playerID, info.Password, info.PlayerName)
	c.Start()

	view.Start()

	err = <-view.Done
	if err != nil {
		logrus.Fatal(err)
	}

	//for {
	//	logrus.Info("hello world!")
	//	time.Sleep(20 * time.Second)
	//}
}
