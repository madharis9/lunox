package main

func main() {
	// sm, err := statemachine.NewFiniteState("booking", statemachine.BookingPlacedEvent)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("1:" + sm.Fsm.Current())

	// currentState := sm.Fsm.Current()
	// err = sm.ChangeState("cancel-booking")
	// if err != nil {
	// 	fmt.Println("err here ", err)
	// }
	// dest := sm.Fsm.Current()
	// if currentState == dest {
	// 	fmt.Println("error")
	// }

	// fmt.Println("2:" + sm.Fsm.Current())

	// fsm := fsm.NewFSM(
	// 	"idle",
	// 	fsm.Events{
	// 		{Name: "scan", Src: []string{"idle"}, Dst: "scanning"},
	// 		{Name: "working", Src: []string{"scanning"}, Dst: "scanning"},
	// 		{Name: "situation", Src: []string{"scanning"}, Dst: "scanning"},
	// 		{Name: "situation", Src: []string{"idle"}, Dst: "idle"},
	// 		{Name: "finish", Src: []string{"scanning"}, Dst: "idle"},
	// 	},
	// 	fsm.Callbacks{
	// 		"scan": func(_ context.Context, e *fsm.Event) {
	// 			fmt.Println("after_scan: " + e.FSM.Current())
	// 		},
	// 		"working": func(_ context.Context, e *fsm.Event) {
	// 			fmt.Println("working: " + e.FSM.Current())
	// 		},
	// 		"situation": func(_ context.Context, e *fsm.Event) {
	// 			fmt.Println("situation: " + e.FSM.Current())
	// 		},
	// 		"finish": func(_ context.Context, e *fsm.Event) {
	// 			fmt.Println("finish: " + e.FSM.Current())
	// 		},
	// 	},
	// )

	// fmt.Println(fsm.Current())

	// err = fsm.Event(context.Background(), "scan")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("1:" + fsm.Current())

	// err = fsm.Event(context.Background(), "working")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("2:" + fsm.Current())

	// err = fsm.Event(context.Background(), "situation")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("3:" + fsm.Current())

	// err = fsm.Event(context.Background(), "finish")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("4:" + fsm.Current())

}
