package main

// 20161003
// 20200706 solventado el problema de la gorutina (hurra!)
// Primitiva ver 0.5.0

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gotk3/gotk3/gdk"
	// "github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var mk1a = "<span foreground=\"black\" background=\"white\" size=\"x-large\"><tt>"
var mk1b = "</tt></span>"
var mk2a = "<span foreground=\"black\" background=\"yellow\" size=\"x-large\"><tt>"
var mk2b = "</tt></span>"
var mk21a = "<span foreground=\"black\" background=\"yellow\" size=\"x-large\"><b><tt>"
var mk21b = "</tt></b></span>"
var mk3a = "<span foreground=\"yellow\" background=\"green\" size=\"x-large\"><tt>"
var mk3b = "</tt></span>"

type Etiquete struct {
	Sel bool
	Lbl *gtk.Label
	Nme string
	Nmx string
}

var etiqs = []Etiquete{}
var lista []string

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	lista = append(lista, "")
	for i := 1; i < 50; i++ {
		lista = append(lista, strconv.Itoa(i))
	}

	gtk.Init(nil)
	win := window()
	win.ShowAll()
	gtk.Main()

}

type Global struct {
	playA  *gtk.Label
	playB  *gtk.Label
	repeat *gtk.Label
	resul  *gtk.Label
	play   *gtk.ToggleButton
	hbox1  *gtk.Box
	vbox1  *gtk.Box
	veces  *gtk.Scale
	delay  *gtk.Scale
	quit   *gtk.Button
}

var gl Global

func window() *gtk.Window {
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetBorderWidth(8)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetTitle("Primitiva")
	win.SetResizable(false)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.Connect("key-press-event", func(_ *gtk.Window, event *gdk.Event) {
		eventKey := gdk.EventKeyNewFromEvent(event)
		fmt.Println(eventKey.KeyVal(), eventKey.State())
	})

	mbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4)
	mlabel, _ := gtk.LabelNew("")
	mlabel.SetMarkup("<span foreground=\"blue\" size=\"x-large\">Primitiva</span>")
	mbox.Add(mlabel)

	primi := primiGrid()
	mbox.Add(primi)

	win.Add(mbox)
	return win
}

func primiGrid() *gtk.Box {
	ind := 0
	count := 0
	sel := []int{}
	fac := 10

	var clear, rand *gtk.Button
	var playing = true

	listax := sameLenght(lista, 1, "r")
	for i, li := range lista {
		var e Etiquete
		e.Lbl, _ = gtk.LabelNew("")
		e.Nme = li
		e.Nmx = listax[i]
		etiqs = append(etiqs, e)
	}

	playOne := func() {
		for _, s := range sel {
			etiqs[s].Lbl.SetMarkup(mk1a + etiqs[s].Nmx + mk1b)
		}
		sel = randNums(6, 49)
		count = len(sel)
		for _, s := range sel {
			etiqs[s].Lbl.SetMarkup(mk3a + etiqs[s].Nmx + mk3b)
		}
	}

	playMulti := func(times, del int) {
		playOneB := func() {
			// for _, s := range sel {
			// 	glib.IdleAdd(markupOne, s)
			// }

			sel = randNums(6, 49)
			count = len(sel)

			// for _, s := range sel {
			// 	glib.IdleAdd(markupTwo, s)
			// }
		}

		for i := 0; i < times; i++ {
			if playing == false {
				playing = true
				break
			}

			// glib.IdleAdd(setPlayA, "  -  -  -  -  -  ")
			// glib.IdleAdd(setPlayB, "  -  -  -  -  -  ")
			// glib.IdleAdd(setRepeat, strconv.Itoa(i+1))
			time.Sleep(time.Duration(del) * time.Millisecond)

			playOneB()

			var resuA = " - "
			var resuB = " - "

			for _, s := range sel {
				resuA += strconv.Itoa(s) + " - "
			}

			// glib.IdleAdd(setPlayA, resuA)
			time.Sleep(time.Duration(del) * time.Millisecond)

			playOneB()

			for _, s := range sel {
				resuB += strconv.Itoa(s) + " - "
			}

			// glib.IdleAdd(setPlayB, resuB)

			if resuA == resuB {
				// glib.IdleAdd(setResul, "ACIERTO")
				break

			} else {
				// glib.IdleAdd(setResul, "FALLO")
			}

			time.Sleep(time.Duration(del) * time.Millisecond)
		}

		// glib.IdleAdd(finalSets)
	}

	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	clear, _ = gtk.ButtonNewWithMnemonic("_Borrar")
	clear.SetCanFocus(false)
	clear.Connect("clicked", func() {
		count = 0
		for _, s := range sel {
			etiqs[s].Lbl.SetMarkup(mk1a + etiqs[s].Nmx + mk1b)
		}
		sel = []int{}
	})

	rand, _ = gtk.ButtonNewWithMnemonic("_Auto")
	rand.SetCanFocus(false)
	rand.Connect("clicked", func() {
		playOne()
	})

	gl.play, _ = gtk.ToggleButtonNewWithMnemonic("_Jugar")
	gl.play.SetCanFocus(false)
	gl.play.Connect("toggled", func() {
		if gl.play.GetActive() == false {
			playing = false
		} else {
			playing = true
			gl.play.SetLabel("_Parar")
			gl.hbox1.Hide()
			gl.veces.Hide()
			gl.delay.Hide()
			gl.vbox1.Show()
			gl.quit.Hide()
			times := int(gl.veces.GetValue())
			del := int(gl.delay.GetValue())
			go playMulti(times, del)
			return
		}
	})

	gl.quit, _ = gtk.ButtonNewWithMnemonic("_Salir")
	gl.quit.SetCanFocus(false)
	gl.quit.SetMarginTop(12)
	gl.quit.Connect("clicked", func() {
		gtk.MainQuit()
	})

	gl.playA, _ = gtk.LabelNew("  -  -  -  -  -  ")
	gl.playA.SetMarginTop(8)
	gl.playB, _ = gtk.LabelNew("  -  -  -  -  -  ")
	gl.playB.SetMarginTop(8)
	gl.resul, _ = gtk.LabelNew("-----")
	gl.resul.SetMarginTop(8)
	gl.repeat, _ = gtk.LabelNew("0")
	gl.repeat.SetMarginTop(8)

	gl.veces, _ = gtk.ScaleNewWithRange(gtk.ORIENTATION_HORIZONTAL, 1, 100, 1)
	gl.veces.SetMarginTop(8)
	gl.veces.SetMarginStart(48)
	gl.veces.SetMarginEnd(48)
	gl.delay, _ = gtk.ScaleNewWithRange(gtk.ORIENTATION_HORIZONTAL, 20, 200, 1)
	gl.delay.SetMarginTop(8)
	gl.delay.SetMarginStart(48)
	gl.delay.SetMarginEnd(48)

	grid, _ := gtk.GridNew()
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)
	grid.SetMarginTop(6)
	grid.SetMarginBottom(6)
	grid.SetMarginStart(6)
	grid.SetRowHomogeneous(true)
	grid.SetRowSpacing(6)

	i, p := 0, 0
	for i < len(etiqs) {
		m := 0
		k := len(etiqs) - p*fac
		if k < fac {
			m = fac - k
			fac = k
		}

		for j := 0; j < fac; j++ {
			pos := i + j
			eti := etiqs[pos].Lbl
			eti.SetMarkup(mk1a + etiqs[pos].Nmx + mk1b)

			ebox, _ := gtk.EventBoxNew()
			ebox.SetVisibleWindow(false)
			ebox.SetMarginEnd(6)
			ebox.SetName(strconv.Itoa(pos))
			ebox.Add(eti)
			grid.Attach(ebox, i, j, 1, 1)

			ebox.Connect("enter-notify-event", func() {
				if count < 6 {
					nm, _ := ebox.GetName()
					nami, _ := strconv.Atoi(nm)
					for _, s := range sel {
						if s == nami {
							return
						}
					}
					etiqs[nami].Lbl.SetMarkup(mk21a + etiqs[nami].Nmx + mk21b)
				}
			})
			ebox.Connect("leave-notify-event", func() {
				if count < 6 {
					nm, _ := ebox.GetName()
					nami, _ := strconv.Atoi(nm)
					for _, s := range sel {
						if s == nami {
							return
						}
					}
					etiqs[nami].Lbl.SetMarkup(mk1a + etiqs[nami].Nmx + mk1b)
				}
			})

			ebox.Connect("button-press-event", func() {
				if count < 6 {
					Nm, _ := ebox.GetName()
					nami, _ := strconv.Atoi(Nm)
					if nami < 1 {
						return
					}
					for i := 0; i < len(sel); i++ {
						if sel[i] == nami {
							etiqs[nami].Lbl.SetMarkup(mk2a + etiqs[nami].Nmx + mk2b)
							return
						}
					}
					etiqs[nami].Lbl.SetMarkup(mk2a + etiqs[nami].Nmx + mk2b)
					count++
					sel = append(sel, nami)
					ind = nami
					clicked(sel)
				}
				if count == 6 {
					for _, s := range sel {
						etiqs[s].Lbl.SetMarkup(mk3a + etiqs[s].Nmx + mk3b)
					}
					count++
				}
			})
		}
		for n := 0; n < m; n++ {
			break
		}
		i += fac
		p++
	}

	frame, _ := gtk.FrameNew("")
	frame.SetMarginBottom(12)
	frame.Add(grid)
	vbox.Add(frame)

	gl.hbox1, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	gl.hbox1.SetHomogeneous(true)
	gl.hbox1.Add(clear)
	gl.hbox1.Add(rand)

	gl.vbox1, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	gl.vbox1.SetHomogeneous(true)
	gl.vbox1.Add(gl.repeat)
	gl.vbox1.Add(gl.playA)
	gl.vbox1.Add(gl.playB)
	gl.vbox1.Add(gl.resul)

	vbox.Add(gl.hbox1)
	vbox.Add(gl.play)
	vbox.Add(gl.veces)
	vbox.Add(gl.delay)
	vbox.Add(gl.vbox1)
	vbox.Add(gl.quit)

	return vbox
}

func sameLenght(lista []string, min int, adj string) []string {
	maxl := 0
	var listax = []string{}
	if min < 0 {
		min = 0
	}

	for _, l := range lista {
		lenl := len(l)
		if lenl > maxl {
			maxl = lenl
		}
	}
	for _, l := range lista {
		var b int
		var as, bs string
		totl := maxl + min*2
		totw := totl - len(l)
		b = totw - min
		if adj == "l" {
			as = strings.Repeat(" ", min)
			bs = strings.Repeat(" ", b)
		} else {
			as = strings.Repeat(" ", b)
			bs = strings.Repeat(" ", min)
		}
		lx := as + l + bs
		listax = append(listax, lx)
	}

	return listax
}

func clicked(sel []int) {
	sort.Ints(sel)
	fmt.Println(sel)
}

func randNums(max, rang int) []int {
	var num int
	var nums []int

	isIn := func(num int) bool {
		for _, n := range nums {
			if num == n {
				return true
			}
		}
		return false
	}
	for i := 0; i < max; i++ {
		num = rand.Intn(rang) + 1
		is := isIn(num)
		for is == true {
			num = rand.Intn(max) + 1
			is = isIn(num)
		}
		nums = append(nums, num)
	}

	sort.Ints(nums)
	return nums
}

func setPlayA(text string) bool {
	gl.playA.SetText(text)
	return false
}

func setPlayB(text string) bool {
	gl.playB.SetText(text)
	return false
}

func setRepeat(text string) bool {
	gl.repeat.SetText(text)
	return false
}

func setResul(text string) bool {
	gl.resul.SetText(text)
	return false
}

func markupOne(s int) bool {
	etiqs[s].Lbl.SetMarkup(mk1a + etiqs[s].Nmx + mk1b)
	return false
}

func markupTwo(s int) bool {
	etiqs[s].Lbl.SetMarkup(mk3a + etiqs[s].Nmx + mk3b)
	return false
}

func finalSets() bool {
	gl.play.SetLabel("_Jugar")
	gl.hbox1.Show()
	gl.veces.Show()
	gl.delay.Show()
	gl.vbox1.Hide()
	gl.quit.Show()

	return false
}
