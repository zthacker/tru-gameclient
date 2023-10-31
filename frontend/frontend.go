package frontend

import (
	"gameclient/actions"
	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
	"github.com/rivo/tview"
	"time"
)

const (
	backgroundColor = tcell.Color234
	textColor       = tcell.ColorWhite
	playerColor     = tcell.ColorWhite
	wallColor       = tcell.Color24
	laserColor      = tcell.ColorRed
	drawFrequency   = 17 * time.Millisecond
)

type View struct {
	App           *tview.Application
	CurrentPlayer uuid.UUID
	pages         *tview.Pages
	drawCallbacks []func()
	viewPort      tview.Primitive
	actionChannel chan actions.Action
	Done          chan error
}

func NewView(actionChannel chan actions.Action) *View {
	app := tview.NewApplication()
	pages := tview.NewPages()
	view := &View{
		App:           app,
		pages:         pages,
		drawCallbacks: make([]func(), 0),
		actionChannel: actionChannel,
		Done:          make(chan error),
	}

	setupViewPort(view)
	app.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		switch e.Key() {
		case tcell.KeyESC:
			app.SetFocus(view.viewPort)
		case tcell.KeyCtrlQ:
			fallthrough
		case tcell.KeyCtrlC:
			app.Stop()
			select {
			case view.Done <- nil:
			default:
			}
		}
		return e
	})

	app.SetRoot(pages, true)

	return view
}

func (v *View) Start() {
	drawTicker := time.NewTicker(17 * time.Millisecond)
	stop := make(chan bool)
	go func() {
		for {
			for _, callBack := range v.drawCallbacks {
				v.App.QueueUpdate(callBack)
			}
			v.App.Draw()
			<-drawTicker.C
			select {
			case <-stop:
				return
			default:
			}
		}
	}()
	go func() {
		err := v.App.Run()
		stop <- true
		drawTicker.Stop()
		select {
		case v.Done <- err:
		default:

		}
	}()
}

func setupViewPort(view *View) {
	box := tview.NewBox().
		SetBorder(true).
		SetTitle("TR Ultimate").
		SetBackgroundColor(backgroundColor)

	box.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		switch e.Key() {
		case tcell.KeyUp:
			view.actionChannel <- actions.MoveAction{
				Direction: 0,
				ID:        view.CurrentPlayer,
				Created:   time.Now(),
			}
		case tcell.KeyDown:
			view.actionChannel <- actions.MoveAction{
				Direction: 1,
				ID:        view.CurrentPlayer,
				Created:   time.Now(),
			}
		case tcell.KeyLeft:
			view.actionChannel <- actions.MoveAction{
				Direction: 2,
				ID:        view.CurrentPlayer,
				Created:   time.Now(),
			}
		case tcell.KeyRight:
			view.actionChannel <- actions.MoveAction{
				Direction: 3,
				ID:        view.CurrentPlayer,
				Created:   time.Now(),
			}
		}
		return e
	})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(box, 0, 1, true)
	view.pages.AddPage("viewport", flex, true, true)
	view.viewPort = box
}
