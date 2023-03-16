package orderbook

import (
	"testing"
)

// Test_Orderbook_Middleware_Checksum aims to cover the case of a broken
// checksum.
func Test_Orderbook_Middleware_Checksum(t *testing.T) {
	var ord *Orderbook
	{
		ord = New()
	}

	for i, x := range testdatac() {
		err := ord.Middleware(x)
		if i == 7 {
			if err == nil {
				t.Fatal("expected update message with index 7 to cause a checksum error")
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

// Test_Orderbook_Middleware_Failure aims to cover the whole process of order
// book management using the failure testdata. The order book middleware
// contains the glue code for processing an update message provided by the
// Kraken websocket. There we handle the initial snapshot message, consecutive
// update messages, and verify the provided checksum against our internal order
// book state.
func Test_Orderbook_Middleware_Failure(t *testing.T) {
	var ord *Orderbook
	{
		ord = New()
	}

	for _, x := range testdataf() {
		err := ord.Middleware(x)
		if err != nil {
			t.Fatal(err)
		}
	}
}

// Test_Orderbook_Middleware_Success aims to cover the whole process of order
// book management using the success testdata. The order book middleware contains
// the glue code for processing an update message provided by the Kraken
// websocket. There we handle the initial snapshot message, consecutive update
// messages, and verify the provided checksum against our internal order book
// state.
func Test_Orderbook_Middleware_Success(t *testing.T) {
	var ord *Orderbook
	{
		ord = New()
	}

	for _, x := range testdatas() {
		err := ord.Middleware(x)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func Benchmark_Orderbook_Middleware(b *testing.B) {
	var ord *Orderbook
	{
		ord = New()
	}

	for i := 0; i < b.N; i++ {
		for _, x := range testdatas()[:100] {
			err := ord.Middleware(x)
			if err != nil {
				panic(err)
			}
		}
	}
}

// testdatac returns raw websocket data similar to testdataf, with the
// difference of a deliberate faulty checksum injected.
func testdatac() []Response {
	return []Response{
		{Asks: []Object{{Price: "1289.60000", Volume: "0.01551181", Time: "1669985812.099879", Republish: false}, {Price: "1289.62000", Volume: "0.18295959", Time: "1669985816.353101", Republish: false}, {Price: "1289.65000", Volume: "0.01551109", Time: "1669985808.813785", Republish: false}, {Price: "1289.73000", Volume: "0.01551036", Time: "1669985826.675521", Republish: false}, {Price: "1289.74000", Volume: "0.37464381", Time: "1669985745.572128", Republish: false}, {Price: "1289.79000", Volume: "0.01550940", Time: "1669985808.824867", Republish: false}, {Price: "1289.85000", Volume: "0.01550868", Time: "1669985817.476325", Republish: false}, {Price: "1289.92000", Volume: "0.01550796", Time: "1669985763.197097", Republish: false}, {Price: "1289.98000", Volume: "0.15508790", Time: "1669985824.325921", Republish: false}, {Price: "1289.99000", Volume: "4.33660106", Time: "1669985825.005016", Republish: false}}, Bids: []Object{{Price: "1289.59000", Volume: "34.36979734", Time: "1669985827.470806", Republish: false}, {Price: "1289.55000", Volume: "2.67347008", Time: "1669985816.339043", Republish: false}, {Price: "1289.54000", Volume: "3.66737696", Time: "1669985815.855607", Republish: false}, {Price: "1289.47000", Volume: "4.71241458", Time: "1669985817.454554", Republish: false}, {Price: "1289.45000", Volume: "58.16433363", Time: "1669985815.083028", Republish: false}, {Price: "1289.36000", Volume: "4.05646690", Time: "1669985824.891917", Republish: false}, {Price: "1289.35000", Volume: "3.85602979", Time: "1669985824.846599", Republish: false}, {Price: "1289.34000", Volume: "7.95246062", Time: "1669985826.577713", Republish: false}, {Price: "1289.33000", Volume: "77.55916097", Time: "1669985824.426763", Republish: false}, {Price: "1289.31000", Volume: "2.38361137", Time: "1669985825.018397", Republish: false}}, CheckSum: "", IsSnapshot: true},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.54000", Volume: "81.21394699", Time: "1669985828.431310", Republish: false}}, CheckSum: "820124158", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.36000", Volume: "11.80699150", Time: "1669985828.456247", Republish: false}}, CheckSum: "1028755919", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.59000", Volume: "34.39764779", Time: "1669985828.719198", Republish: false}}, CheckSum: "1339011945", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.59000", Volume: "34.36979734", Time: "1669985829.189713", Republish: false}}, CheckSum: "1028755919", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.36000", Volume: "15.47233677", Time: "1669985830.318128", Republish: false}}, CheckSum: "483980390", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.37000", Volume: "23.25752868", Time: "1669985830.585508", Republish: false}}, CheckSum: "1050857887", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.59000", Volume: "45.99782697", Time: "1669985830.835409", Republish: false}}, CheckSum: "486151523", IsSnapshot: false}, // checksum manipulated (incremented by 1)
	}
}

// testdataf returns raw websocket data for which the checksum calculation could
// initially NOT be done without issues.
func testdataf() []Response {
	return []Response{
		{Asks: []Object{{Price: "1289.60000", Volume: "0.01551181", Time: "1669985812.099879", Republish: false}, {Price: "1289.62000", Volume: "0.18295959", Time: "1669985816.353101", Republish: false}, {Price: "1289.65000", Volume: "0.01551109", Time: "1669985808.813785", Republish: false}, {Price: "1289.73000", Volume: "0.01551036", Time: "1669985826.675521", Republish: false}, {Price: "1289.74000", Volume: "0.37464381", Time: "1669985745.572128", Republish: false}, {Price: "1289.79000", Volume: "0.01550940", Time: "1669985808.824867", Republish: false}, {Price: "1289.85000", Volume: "0.01550868", Time: "1669985817.476325", Republish: false}, {Price: "1289.92000", Volume: "0.01550796", Time: "1669985763.197097", Republish: false}, {Price: "1289.98000", Volume: "0.15508790", Time: "1669985824.325921", Republish: false}, {Price: "1289.99000", Volume: "4.33660106", Time: "1669985825.005016", Republish: false}}, Bids: []Object{{Price: "1289.59000", Volume: "34.36979734", Time: "1669985827.470806", Republish: false}, {Price: "1289.55000", Volume: "2.67347008", Time: "1669985816.339043", Republish: false}, {Price: "1289.54000", Volume: "3.66737696", Time: "1669985815.855607", Republish: false}, {Price: "1289.47000", Volume: "4.71241458", Time: "1669985817.454554", Republish: false}, {Price: "1289.45000", Volume: "58.16433363", Time: "1669985815.083028", Republish: false}, {Price: "1289.36000", Volume: "4.05646690", Time: "1669985824.891917", Republish: false}, {Price: "1289.35000", Volume: "3.85602979", Time: "1669985824.846599", Republish: false}, {Price: "1289.34000", Volume: "7.95246062", Time: "1669985826.577713", Republish: false}, {Price: "1289.33000", Volume: "77.55916097", Time: "1669985824.426763", Republish: false}, {Price: "1289.31000", Volume: "2.38361137", Time: "1669985825.018397", Republish: false}}, CheckSum: "", IsSnapshot: true},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.54000", Volume: "81.21394699", Time: "1669985828.431310", Republish: false}}, CheckSum: "820124158", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.36000", Volume: "11.80699150", Time: "1669985828.456247", Republish: false}}, CheckSum: "1028755919", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.59000", Volume: "34.39764779", Time: "1669985828.719198", Republish: false}}, CheckSum: "1339011945", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.59000", Volume: "34.36979734", Time: "1669985829.189713", Republish: false}}, CheckSum: "1028755919", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.36000", Volume: "15.47233677", Time: "1669985830.318128", Republish: false}}, CheckSum: "483980390", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.37000", Volume: "23.25752868", Time: "1669985830.585508", Republish: false}}, CheckSum: "1050857887", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.59000", Volume: "45.99782697", Time: "1669985830.835409", Republish: false}}, CheckSum: "486151522", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.59000", Volume: "46.01531986", Time: "1669985830.862066", Republish: false}}, CheckSum: "1238184002", IsSnapshot: false},
		{Asks: []Object{{Price: "1289.99000", Volume: "4.49167543", Time: "1669985830.885237", Republish: false}}, Bids: []Object(nil), CheckSum: "4210909853", IsSnapshot: false},
		{Asks: []Object{{Price: "1289.98000", Volume: "0.00000000", Time: "1669985830.885322", Republish: false}, {Price: "1290.00000", Volume: "0.01000000", Time: "1669985822.198976", Republish: true}}, Bids: []Object(nil), CheckSum: "512919999", IsSnapshot: false},
		{Asks: []Object{{Price: "1289.99000", Volume: "4.33660106", Time: "1669985830.918593", Republish: false}}, Bids: []Object(nil), CheckSum: "2479652596", IsSnapshot: false},
		{Asks: []Object{{Price: "1289.98000", Volume: "0.15507437", Time: "1669985830.918623", Republish: false}}, Bids: []Object(nil), CheckSum: "2702833408", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.47000", Volume: "7.01192818", Time: "1669985831.453598", Republish: false}}, CheckSum: "3577332032", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1289.37000", Volume: "0.00000000", Time: "1669985832.326406", Republish: false}, {Price: "1289.27000", Volume: "0.01551265", Time: "1669985732.456502", Republish: true}}, CheckSum: "3511863351", IsSnapshot: false},
	}
}

// testdatas returns raw websocket data for which the checksum calculation could
// initially be done without issues.
func testdatas() []Response {
	return []Response{
		{Asks: []Object{{Price: "1272.70000", Volume: "1.00000000", Time: "1669902400.950657", Republish: false}, {Price: "1272.71000", Volume: "2.39742752", Time: "1669902400.924689", Republish: false}, {Price: "1272.72000", Volume: "3.72191238", Time: "1669902400.751811", Republish: false}, {Price: "1272.74000", Volume: "5.50000000", Time: "1669902400.561752", Republish: false}, {Price: "1272.77000", Volume: "4.05813845", Time: "1669902400.783276", Republish: false}, {Price: "1272.78000", Volume: "25.58235843", Time: "1669902400.589057", Republish: false}, {Price: "1272.79000", Volume: "5.00000000", Time: "1669902400.593059", Republish: false}, {Price: "1272.80000", Volume: "2.99895264", Time: "1669902400.133970", Republish: false}, {Price: "1272.85000", Volume: "2.24585184", Time: "1669902400.018159", Republish: false}, {Price: "1272.86000", Volume: "1.00000000", Time: "1669902399.934519", Republish: false}}, Bids: []Object{{Price: "1271.80000", Volume: "4.03174746", Time: "1669902400.824471", Republish: false}, {Price: "1271.79000", Volume: "2.26656312", Time: "1669902391.019095", Republish: false}, {Price: "1271.60000", Volume: "0.65439574", Time: "1669902400.732617", Republish: false}, {Price: "1271.59000", Volume: "26.08235843", Time: "1669902400.588707", Republish: false}, {Price: "1271.58000", Volume: "5.50000000", Time: "1669902399.728626", Republish: false}, {Price: "1271.57000", Volume: "3.87724786", Time: "1669902399.586119", Republish: false}, {Price: "1271.56000", Volume: "3.89057615", Time: "1669902396.164743", Republish: false}, {Price: "1271.54000", Volume: "4.74487479", Time: "1669902395.334591", Republish: false}, {Price: "1271.49000", Volume: "58.98572653", Time: "1669902394.921835", Republish: false}, {Price: "1271.40000", Volume: "1.93804942", Time: "1669902389.029716", Republish: false}}, CheckSum: "", IsSnapshot: true},
		{Asks: []Object{{Price: "1272.74000", Volume: "0.00000000", Time: "1669902401.097616", Republish: false}, {Price: "1272.95000", Volume: "6.22500000", Time: "1669902400.976119", Republish: true}}, Bids: []Object(nil), CheckSum: "3809965286", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.70000", Volume: "6.50000000", Time: "1669902401.097779", Republish: false}}, Bids: []Object(nil), CheckSum: "701627836", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.86000", Volume: "0.00000000", Time: "1669902401.125255", Republish: false}, {Price: "1272.98000", Volume: "12.17615359", Time: "1669902397.735301", Republish: true}}, Bids: []Object(nil), CheckSum: "2495027285", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.60000", Volume: "0.00000000", Time: "1669902401.394152", Republish: false}, {Price: "1271.04000", Volume: "7.76200000", Time: "1669902400.521452", Republish: true}}, CheckSum: "2762515866", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.81000", Volume: "0.55238494", Time: "1669902401.444091", Republish: false}}, CheckSum: "3196041703", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.69000", Volume: "5.00000000", Time: "1669902401.470356", Republish: false}}, Bids: []Object(nil), CheckSum: "3987179842", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.60000", Volume: "2.50000000", Time: "1669902401.490049", Republish: false}}, CheckSum: "497887866", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.70000", Volume: "1.00000000", Time: "1669902401.505545", Republish: false}}, Bids: []Object(nil), CheckSum: "3482256474", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.68000", Volume: "5.50000000", Time: "1669902401.505628", Republish: false}}, Bids: []Object(nil), CheckSum: "218861191", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.81000", Volume: "8.31438494", Time: "1669902401.514437", Republish: false}}, CheckSum: "3699944351", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.79000", Volume: "0.00000000", Time: "1669902401.516647", Republish: false}, {Price: "1272.95000", Volume: "6.22500000", Time: "1669902400.976119", Republish: true}}, Bids: []Object(nil), CheckSum: "182075662", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.81000", Volume: "0.55238494", Time: "1669902401.535037", Republish: false}}, CheckSum: "2302867239", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.68000", Volume: "0.00000000", Time: "1669902401.614967", Republish: false}, {Price: "1272.98000", Volume: "12.17615359", Time: "1669902397.735301", Republish: true}}, Bids: []Object(nil), CheckSum: "3924650082", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.72000", Volume: "0.00000000", Time: "1669902401.633597", Republish: false}, {Price: "1273.21000", Volume: "1.36697687", Time: "1669902398.285287", Republish: true}, {Price: "1272.67000", Volume: "3.72191238", Time: "1669902401.633619", Republish: false}}, Bids: []Object(nil), CheckSum: "230589470", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.69000", Volume: "0.00000000", Time: "1669902401.760977", Republish: false}, {Price: "1273.21000", Volume: "1.36697687", Time: "1669902398.285287", Republish: true}}, Bids: []Object(nil), CheckSum: "4141684802", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.12000", Volume: "1.64520929", Time: "1669902401.778252", Republish: false}}, Bids: []Object(nil), CheckSum: "142788835", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.70000", Volume: "0.00000000", Time: "1669902401.778774", Republish: false}, {Price: "1273.21000", Volume: "1.36697687", Time: "1669902398.285287", Republish: true}}, Bids: []Object(nil), CheckSum: "3884337564", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.95000", Volume: "0.00000000", Time: "1669902401.779781", Republish: false}, {Price: "1273.30000", Volume: "8.61283478", Time: "1669902400.250177", Republish: true}}, Bids: []Object(nil), CheckSum: "3564328412", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.78000", Volume: "23.58235843", Time: "1669902401.790177", Republish: false}}, Bids: []Object(nil), CheckSum: "3248375459", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.12000", Volume: "0.00000000", Time: "1669902401.793214", Republish: false}, {Price: "1273.31000", Volume: "58.90177825", Time: "1669902399.308586", Republish: true}}, Bids: []Object(nil), CheckSum: "2149605935", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.98000", Volume: "14.17615359", Time: "1669902401.815282", Republish: false}}, Bids: []Object(nil), CheckSum: "1670613157", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.78000", Volume: "0.00000000", Time: "1669902401.827808", Republish: false}, {Price: "1273.35000", Volume: "1.95524348", Time: "1669902401.295910", Republish: true}}, Bids: []Object(nil), CheckSum: "1160059663", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.15000", Volume: "3.43406606", Time: "1669902401.853188", Republish: false}}, Bids: []Object(nil), CheckSum: "1605249319", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.10000", Volume: "6.22500000", Time: "1669902401.859153", Republish: false}}, Bids: []Object(nil), CheckSum: "1789159723", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.15000", Volume: "0.00000000", Time: "1669902401.881787", Republish: false}, {Price: "1273.31000", Volume: "58.90177825", Time: "1669902399.308586", Republish: true}}, Bids: []Object(nil), CheckSum: "2397248057", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.80000", Volume: "3.55090016", Time: "1669902401.894332", Republish: false}}, CheckSum: "2831340374", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.16000", Volume: "7.19756861", Time: "1669902402.009018", Republish: false}}, Bids: []Object(nil), CheckSum: "2909371969", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.82000", Volume: "1.00000000", Time: "1669902402.036364", Republish: false}}, CheckSum: "2465187246", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.81000", Volume: "0.00000000", Time: "1669902402.128570", Republish: false}, {Price: "1271.49000", Volume: "58.98572653", Time: "1669902394.921835", Republish: true}}, CheckSum: "340664821", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.82000", Volume: "0.00000000", Time: "1669902402.155502", Republish: false}, {Price: "1271.46000", Volume: "12.19070989", Time: "1669902401.584362", Republish: true}}, CheckSum: "3572424830", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.81000", Volume: "0.59361121", Time: "1669902402.236524", Republish: false}}, CheckSum: "2259984358", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.82000", Volume: "1.00000000", Time: "1669902402.261617", Republish: false}}, CheckSum: "1994596045", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.81000", Volume: "1.07445851", Time: "1669902402.306386", Republish: false}}, CheckSum: "3424699854", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.81000", Volume: "0.48084730", Time: "1669902402.327429", Republish: false}}, CheckSum: "1834967880", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.55000", Volume: "0.55104218", Time: "1669902402.329155", Republish: false}}, CheckSum: "1239756070", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.82000", Volume: "3.29298000", Time: "1669902402.339247", Republish: false}}, CheckSum: "239395077", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.16000", Volume: "0.00000000", Time: "1669902402.368631", Republish: false}, {Price: "1273.31000", Volume: "58.90177825", Time: "1669902399.308586", Republish: true}}, Bids: []Object(nil), CheckSum: "628981546", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.55000", Volume: "0.00000000", Time: "1669902402.370638", Republish: false}, {Price: "1271.54000", Volume: "4.74487479", Time: "1669902395.334591", Republish: true}}, CheckSum: "1686484785", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.83000", Volume: "6.22500000", Time: "1669902402.396058", Republish: false}}, CheckSum: "3103471294", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.83000", Volume: "6.81945060", Time: "1669902402.431026", Republish: false}}, CheckSum: "1259953504", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.66000", Volume: "2.46042000", Time: "1669902402.449974", Republish: false}}, Bids: []Object(nil), CheckSum: "3683331987", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.84000", Volume: "5.00000000", Time: "1669902402.463296", Republish: false}}, CheckSum: "577362008", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.57000", Volume: "1.37724786", Time: "1669902402.478797", Republish: false}}, CheckSum: "1375071368", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.23000", Volume: "2.50000000", Time: "1669902402.479230", Republish: false}}, CheckSum: "2312451141", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.83000", Volume: "6.22500000", Time: "1669902402.480415", Republish: false}}, CheckSum: "1762623302", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.67000", Volume: "0.00000000", Time: "1669902402.537054", Republish: false}, {Price: "1273.31000", Volume: "58.90177825", Time: "1669902399.308586", Republish: true}, {Price: "1273.29000", Volume: "3.72191238", Time: "1669902402.537072", Republish: false}}, Bids: []Object(nil), CheckSum: "1619027803", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.23000", Volume: "4.33816096", Time: "1669902402.554680", Republish: false}}, CheckSum: "1191498730", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.85000", Volume: "0.62747402", Time: "1669902402.583254", Republish: false}}, CheckSum: "2765411381", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.22000", Volume: "4.80358033", Time: "1669902402.626458", Republish: false}}, Bids: []Object(nil), CheckSum: "3893738411", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.10000", Volume: "0.00000000", Time: "1669902402.674661", Republish: false}, {Price: "1273.30000", Volume: "8.61283478", Time: "1669902400.250177", Republish: true}}, Bids: []Object(nil), CheckSum: "3705585581", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.71000", Volume: "0.00000000", Time: "1669902402.683354", Republish: false}, {Price: "1273.31000", Volume: "58.90177825", Time: "1669902399.308586", Republish: true}}, Bids: []Object(nil), CheckSum: "2298616655", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.22000", Volume: "0.00000000", Time: "1669902402.690237", Republish: false}, {Price: "1273.35000", Volume: "1.95524348", Time: "1669902401.295910", Republish: true}}, Bids: []Object(nil), CheckSum: "3840498644", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.66000", Volume: "0.00000000", Time: "1669902402.691144", Republish: false}, {Price: "1273.70000", Volume: "7.76200000", Time: "1669902401.955846", Republish: true}}, Bids: []Object(nil), CheckSum: "2318276612", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.59000", Volume: "2.50000000", Time: "1669902402.694816", Republish: false}}, CheckSum: "2272095732", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.31000", Volume: "61.89236917", Time: "1669902402.708796", Republish: false}}, Bids: []Object(nil), CheckSum: "3925629667", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.98000", Volume: "12.17615359", Time: "1669902402.710118", Republish: false}}, Bids: []Object(nil), CheckSum: "2265821356", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.70000", Volume: "0.00000000", Time: "1669902402.710812", Republish: false}, {Price: "1273.74000", Volume: "1.72447597", Time: "1669902399.745722", Republish: true}}, Bids: []Object(nil), CheckSum: "4116444970", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.31000", Volume: "58.90177825", Time: "1669902402.723437", Republish: false}}, Bids: []Object(nil), CheckSum: "2613755965", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.48000", Volume: "5.50000000", Time: "1669902402.725132", Republish: false}}, CheckSum: "3751589675", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.20000", Volume: "2.00000000", Time: "1669902402.735545", Republish: false}}, Bids: []Object(nil), CheckSum: "2274496573", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.31000", Volume: "0.00000000", Time: "1669902402.739400", Republish: false}, {Price: "1273.74000", Volume: "1.72447597", Time: "1669902399.745722", Republish: true}}, Bids: []Object(nil), CheckSum: "1128073416", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.87000", Volume: "12.18678009", Time: "1669902402.740098", Republish: false}}, CheckSum: "1162891626", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.85000", Volume: "8.38947402", Time: "1669902402.741951", Republish: false}}, CheckSum: "3922081167", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.77000", Volume: "4.01810502", Time: "1669902402.745054", Republish: false}}, Bids: []Object(nil), CheckSum: "3224738565", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.80000", Volume: "0.00000000", Time: "1669902402.748513", Republish: false}, {Price: "1273.88000", Volume: "2.44900000", Time: "1669902397.993331", Republish: true}}, Bids: []Object(nil), CheckSum: "1686996327", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.29000", Volume: "0.00000000", Time: "1669902402.763599", Republish: false}, {Price: "1274.11000", Volume: "9.55813836", Time: "1669902399.743740", Republish: true}, {Price: "1274.10000", Volume: "3.72191238", Time: "1669902402.763613", Republish: false}}, Bids: []Object(nil), CheckSum: "656932699", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.77000", Volume: "0.00000000", Time: "1669902402.772650", Republish: false}, {Price: "1274.11000", Volume: "9.55813836", Time: "1669902399.743740", Republish: true}}, Bids: []Object(nil), CheckSum: "3918037542", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.87000", Volume: "0.00000000", Time: "1669902402.779145", Republish: false}, {Price: "1271.60000", Volume: "2.50000000", Time: "1669902401.490049", Republish: true}}, CheckSum: "3421338738", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.85000", Volume: "0.00000000", Time: "1669902402.782710", Republish: false}, {Price: "1274.13000", Volume: "1.70743124", Time: "1669902401.295435", Republish: true}}, Bids: []Object(nil), CheckSum: "3658207640", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.21000", Volume: "0.00000000", Time: "1669902402.787513", Republish: false}, {Price: "1274.21000", Volume: "12.16439990", Time: "1669902397.187795", Republish: true}, {Price: "1273.68000", Volume: "1.36697687", Time: "1669902402.787532", Republish: false}}, Bids: []Object(nil), CheckSum: "908802850", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.63000", Volume: "0.61877242", Time: "1669902402.810909", Republish: false}}, Bids: []Object(nil), CheckSum: "2310837574", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.85000", Volume: "7.76200000", Time: "1669902402.816974", Republish: false}}, CheckSum: "1172203586", IsSnapshot: false},
		{Asks: []Object{{Price: "1272.98000", Volume: "0.00000000", Time: "1669902402.820974", Republish: false}, {Price: "1274.13000", Volume: "1.70743124", Time: "1669902401.295435", Republish: true}}, Bids: []Object(nil), CheckSum: "246901347", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.30000", Volume: "0.00000000", Time: "1669902402.847443", Republish: false}, {Price: "1274.21000", Volume: "12.16439990", Time: "1669902397.187795", Republish: true}}, Bids: []Object(nil), CheckSum: "2220606401", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.63000", Volume: "0.00000000", Time: "1669902402.863888", Republish: false}, {Price: "1274.22000", Volume: "52.67026943", Time: "1669902391.729762", Republish: true}}, Bids: []Object(nil), CheckSum: "2049571307", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.35000", Volume: "0.00000000", Time: "1669902402.868522", Republish: false}, {Price: "1274.52000", Volume: "0.08115090", Time: "1669902401.296531", Republish: true}}, Bids: []Object(nil), CheckSum: "2892505404", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.52000", Volume: "0.00000000", Time: "1669902402.868768", Republish: false}, {Price: "1275.00000", Volume: "61.40249992", Time: "1669902401.643491", Republish: true}}, Bids: []Object(nil), CheckSum: "2226869805", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.49000", Volume: "5.00000000", Time: "1669902402.874149", Republish: false}}, CheckSum: "1345090098", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.24000", Volume: "1.37724786", Time: "1669902402.885182", Republish: false}}, CheckSum: "714461186", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.09000", Volume: "8.28177245", Time: "1669902402.894020", Republish: false}}, Bids: []Object(nil), CheckSum: "3203770542", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.62000", Volume: "0.55876870", Time: "1669902402.906922", Republish: false}}, Bids: []Object(nil), CheckSum: "639805167", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.24000", Volume: "1.94559666", Time: "1669902402.908000", Republish: false}}, CheckSum: "295645976", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.84000", Volume: "0.00000000", Time: "1669902402.910830", Republish: false}, {Price: "1271.79000", Volume: "2.26656312", Time: "1669902391.019095", Republish: true}}, CheckSum: "2225827", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.24000", Volume: "1.37724786", Time: "1669902402.965942", Republish: false}}, CheckSum: "932995924", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.88000", Volume: "0.00000000", Time: "1669902402.970184", Republish: false}, {Price: "1274.22000", Volume: "52.67026943", Time: "1669902391.729762", Republish: true}}, Bids: []Object(nil), CheckSum: "1939593456", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.36000", Volume: "6.22500000", Time: "1669902402.971573", Republish: false}}, CheckSum: "3620306991", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.97000", Volume: "3.69150373", Time: "1669902402.975720", Republish: false}}, Bids: []Object(nil), CheckSum: "3674365284", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.50000", Volume: "1.00000000", Time: "1669902402.994739", Republish: false}}, CheckSum: "4109189857", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.48000", Volume: "0.00000000", Time: "1669902403.019285", Republish: false}, {Price: "1271.80000", Volume: "3.55090016", Time: "1669902401.894332", Republish: true}}, CheckSum: "3745520847", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.55000", Volume: "5.50000000", Time: "1669902403.019421", Republish: false}}, CheckSum: "3166147381", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.50000", Volume: "1.48013575", Time: "1669902403.021316", Republish: false}}, CheckSum: "104831798", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.93000", Volume: "0.37323852", Time: "1669902403.034392", Republish: false}}, Bids: []Object(nil), CheckSum: "3016620605", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.68000", Volume: "0.00000000", Time: "1669902403.043734", Republish: false}, {Price: "1274.21000", Volume: "12.16439990", Time: "1669902397.187795", Republish: true}, {Price: "1274.09000", Volume: "9.64874932", Time: "1669902403.043750", Republish: false}}, Bids: []Object(nil), CheckSum: "3229743334", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.62000", Volume: "0.00000000", Time: "1669902403.045514", Republish: false}, {Price: "1274.22000", Volume: "52.67026943", Time: "1669902391.729762", Republish: true}}, Bids: []Object(nil), CheckSum: "1655073145", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.50000", Volume: "0.48013575", Time: "1669902403.046091", Republish: false}}, CheckSum: "2463610667", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.83000", Volume: "0.00000000", Time: "1669902403.048594", Republish: false}, {Price: "1271.80000", Volume: "3.55090016", Time: "1669902401.894332", Republish: true}}, CheckSum: "3646556005", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.93000", Volume: "0.00000000", Time: "1669902403.093107", Republish: false}, {Price: "1274.26000", Volume: "7.76200000", Time: "1669902402.984943", Republish: true}}, Bids: []Object(nil), CheckSum: "1076975800", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.02000", Volume: "1.08550462", Time: "1669902403.115849", Republish: false}}, Bids: []Object(nil), CheckSum: "323341838", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.97000", Volume: "0.00000000", Time: "1669902403.146956", Republish: false}, {Price: "1274.26000", Volume: "7.76200000", Time: "1669902402.984943", Republish: true}}, Bids: []Object(nil), CheckSum: "3278637413", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.95000", Volume: "0.81401727", Time: "1669902403.194820", Republish: false}}, Bids: []Object(nil), CheckSum: "2521247944", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.20000", Volume: "1.96073158", Time: "1669902403.301387", Republish: false}}, Bids: []Object(nil), CheckSum: "3514436187", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.86000", Volume: "70.66025269", Time: "1669902403.327022", Republish: false}}, Bids: []Object(nil), CheckSum: "267787388", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.20000", Volume: "0.00000000", Time: "1669902403.327238", Republish: false}, {Price: "1274.22000", Volume: "52.67026943", Time: "1669902391.729762", Republish: true}}, Bids: []Object(nil), CheckSum: "4264497864", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1271.85000", Volume: "0.00000000", Time: "1669902403.345463", Republish: false}, {Price: "1271.79000", Volume: "2.26656312", Time: "1669902391.019095", Republish: true}, {Price: "1272.24000", Volume: "9.13924786", Time: "1669902403.345480", Republish: false}}, CheckSum: "1271561017", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.86000", Volume: "0.00000000", Time: "1669902403.354526", Republish: false}, {Price: "1274.26000", Volume: "7.76200000", Time: "1669902402.984943", Republish: true}}, Bids: []Object(nil), CheckSum: "2711118335", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.95000", Volume: "0.00000000", Time: "1669902403.367666", Republish: false}, {Price: "1275.00000", Volume: "61.40249992", Time: "1669902401.643491", Republish: true}}, Bids: []Object(nil), CheckSum: "1755349846", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.52000", Volume: "1.78523696", Time: "1669902403.373681", Republish: false}}, Bids: []Object(nil), CheckSum: "2676544774", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.02000", Volume: "0.00000000", Time: "1669902403.375576", Republish: false}, {Price: "1274.91000", Volume: "0.08890337", Time: "1669902403.374601", Republish: true}}, Bids: []Object(nil), CheckSum: "3992918849", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.26000", Volume: "0.00000000", Time: "1669902403.402393", Republish: false}, {Price: "1275.00000", Volume: "61.40249992", Time: "1669902401.643491", Republish: true}}, Bids: []Object(nil), CheckSum: "3921173600", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.07000", Volume: "2.50000000", Time: "1669902403.467445", Republish: false}}, Bids: []Object(nil), CheckSum: "3702348287", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.56000", Volume: "2.50000000", Time: "1669902403.468187", Republish: false}}, CheckSum: "3911249711", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.56000", Volume: "7.50000000", Time: "1669902403.493091", Republish: false}}, CheckSum: "3053790852", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.57000", Volume: "2.00000000", Time: "1669902403.494450", Republish: false}}, CheckSum: "3635429917", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.65000", Volume: "58.93178495", Time: "1669902403.509657", Republish: false}}, CheckSum: "777246505", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.73000", Volume: "0.48084730", Time: "1669902403.512614", Republish: false}}, Bids: []Object(nil), CheckSum: "2943418957", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.55000", Volume: "0.00000000", Time: "1669902403.520200", Republish: false}, {Price: "1271.81000", Volume: "0.48084730", Time: "1669902402.327429", Republish: true}}, CheckSum: "907520456", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.61000", Volume: "5.50000000", Time: "1669902403.520773", Republish: false}}, CheckSum: "3796777414", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.39000", Volume: "12.18179960", Time: "1669902403.521893", Republish: false}}, CheckSum: "361916565", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.50000", Volume: "0.00000000", Time: "1669902403.528506", Republish: false}, {Price: "1271.82000", Volume: "3.29298000", Time: "1669902402.339247", Republish: true}}, CheckSum: "577931503", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.56000", Volume: "2.50000000", Time: "1669902403.538887", Republish: false}}, CheckSum: "3655255226", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.11000", Volume: "3.66271550", Time: "1669902403.539339", Republish: false}}, Bids: []Object(nil), CheckSum: "789635857", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.20000", Volume: "1.07915345", Time: "1669902403.540655", Republish: false}}, Bids: []Object(nil), CheckSum: "611348212", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.61000", Volume: "0.00000000", Time: "1669902403.540989", Republish: false}, {Price: "1271.81000", Volume: "0.48084730", Time: "1669902402.327429", Republish: true}}, CheckSum: "3622827517", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.70000", Volume: "5.50000000", Time: "1669902403.541892", Republish: false}}, CheckSum: "1037709251", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.20000", Volume: "0.00000000", Time: "1669902403.555545", Republish: false}, {Price: "1274.52000", Volume: "1.78523696", Time: "1669902403.373681", Republish: true}}, Bids: []Object(nil), CheckSum: "918200870", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.09000", Volume: "8.28177245", Time: "1669902403.558527", Republish: false}, {Price: "1274.20000", Volume: "1.36697687", Time: "1669902403.558542", Republish: false}}, Bids: []Object(nil), CheckSum: "227096367", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.56000", Volume: "26.05579802", Time: "1669902403.587061", Republish: false}}, CheckSum: "4221768964", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.08000", Volume: "23.55579802", Time: "1669902403.587253", Republish: false}}, Bids: []Object(nil), CheckSum: "1247171660", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.11000", Volume: "0.00000000", Time: "1669902403.587351", Republish: false}, {Price: "1274.22000", Volume: "52.67026943", Time: "1669902391.729762", Republish: true}}, Bids: []Object(nil), CheckSum: "4086459020", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.74000", Volume: "0.00000000", Time: "1669902403.621492", Republish: false}, {Price: "1274.52000", Volume: "1.78523696", Time: "1669902403.373681", Republish: true}}, Bids: []Object(nil), CheckSum: "4221566491", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.72000", Volume: "12.16907954", Time: "1669902403.635826", Republish: false}}, Bids: []Object(nil), CheckSum: "1285667198", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.72000", Volume: "10.41214809", Time: "1669902403.691454", Republish: false}}, Bids: []Object(nil), CheckSum: "2224265278", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.21000", Volume: "0.00000000", Time: "1669902403.731516", Republish: false}, {Price: "1274.52000", Volume: "1.78523696", Time: "1669902403.373681", Republish: true}}, Bids: []Object(nil), CheckSum: "2585068708", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.10000", Volume: "0.00000000", Time: "1669902403.744242", Republish: false}, {Price: "1274.91000", Volume: "1.78523696", Time: "1669902403.672370", Republish: true}, {Price: "1274.06000", Volume: "3.72191238", Time: "1669902403.744257", Republish: false}}, Bids: []Object(nil), CheckSum: "2798690300", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.73000", Volume: "0.00000000", Time: "1669902403.815815", Republish: false}, {Price: "1274.91000", Volume: "1.78523696", Time: "1669902403.672370", Republish: true}}, Bids: []Object(nil), CheckSum: "1253873384", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.09000", Volume: "0.00000000", Time: "1669902403.858366", Republish: false}, {Price: "1274.98000", Volume: "3.86114224", Time: "1669902403.667804", Republish: true}}, Bids: []Object(nil), CheckSum: "1433663890", IsSnapshot: false},
		{Asks: []Object{{Price: "1273.72000", Volume: "0.00000000", Time: "1669902403.859702", Republish: false}, {Price: "1274.99000", Volume: "5.51607274", Time: "1669902403.588453", Republish: true}}, Bids: []Object(nil), CheckSum: "753466432", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.71000", Volume: "0.03918122", Time: "1669902403.866661", Republish: false}}, CheckSum: "1602519594", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.31000", Volume: "0.06347304", Time: "1669902403.873702", Republish: false}}, Bids: []Object(nil), CheckSum: "2971454065", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.31000", Volume: "0.00000000", Time: "1669902403.888552", Republish: false}, {Price: "1274.99000", Volume: "5.51607274", Time: "1669902403.588453", Republish: true}}, Bids: []Object(nil), CheckSum: "1602519594", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.08000", Volume: "0.00000000", Time: "1669902403.896122", Republish: false}, {Price: "1275.00000", Volume: "61.40249992", Time: "1669902401.643491", Republish: true}}, Bids: []Object(nil), CheckSum: "34509902", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.40000", Volume: "4.04378535", Time: "1669902403.905724", Republish: false}}, CheckSum: "4060618378", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.97000", Volume: "8.28177245", Time: "1669902403.905939", Republish: false}}, Bids: []Object(nil), CheckSum: "444406350", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.24000", Volume: "1.37724786", Time: "1669902403.947276", Republish: false}, {Price: "1272.71000", Volume: "7.80118122", Time: "1669902403.947292", Republish: false}}, CheckSum: "3868289093", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.71000", Volume: "14.02618122", Time: "1669902403.970673", Republish: false}}, CheckSum: "2635190876", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.70000", Volume: "0.00000000", Time: "1669902403.971042", Republish: false}, {Price: "1272.23000", Volume: "4.33816096", Time: "1669902402.554680", Republish: true}}, CheckSum: "1402916184", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.76000", Volume: "5.50000000", Time: "1669902403.971464", Republish: false}}, CheckSum: "925943905", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.72000", Volume: "1.00000000", Time: "1669902403.973620", Republish: false}}, CheckSum: "2583286310", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.93000", Volume: "2.43100000", Time: "1669902403.994400", Republish: false}}, Bids: []Object(nil), CheckSum: "398565310", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.20000", Volume: "0.00000000", Time: "1669902404.018228", Republish: false}, {Price: "1274.99000", Volume: "5.51607274", Time: "1669902403.588453", Republish: true}, {Price: "1274.90000", Volume: "1.36697687", Time: "1669902404.018249", Republish: false}}, Bids: []Object(nil), CheckSum: "3082810121", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.72000", Volume: "0.00000000", Time: "1669902404.025742", Republish: false}, {Price: "1272.24000", Volume: "1.37724786", Time: "1669902403.947276", Republish: true}}, CheckSum: "427223374", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.36000", Volume: "0.00000000", Time: "1669902404.028413", Republish: false}, {Price: "1272.23000", Volume: "4.33816096", Time: "1669902402.554680", Republish: true}}, CheckSum: "1550946547", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.77000", Volume: "0.48013575", Time: "1669902404.045185", Republish: false}}, CheckSum: "617090878", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.77000", Volume: "2.98013575", Time: "1669902404.048436", Republish: false}}, CheckSum: "1005246551", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.04000", Volume: "2.50000000", Time: "1669902404.048861", Republish: false}}, Bids: []Object(nil), CheckSum: "1138495134", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.78000", Volume: "0.15698956", Time: "1669902404.161609", Republish: false}}, CheckSum: "4096356051", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.03000", Volume: "0.15598391", Time: "1669902404.163677", Republish: false}}, Bids: []Object(nil), CheckSum: "1145028029", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.78000", Volume: "6.38198956", Time: "1669902404.191007", Republish: false}}, CheckSum: "1157252198", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.69000", Volume: "0.05323923", Time: "1669902404.202550", Republish: false}}, Bids: []Object(nil), CheckSum: "1685345724", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.78000", Volume: "6.22500000", Time: "1669902404.214549", Republish: false}}, CheckSum: "3942488501", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.83000", Volume: "5.50000000", Time: "1669902404.215904", Republish: false}}, CheckSum: "922955594", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.79000", Volume: "0.15698956", Time: "1669902404.216665", Republish: false}}, CheckSum: "855523122", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.69000", Volume: "0.00000000", Time: "1669902404.224090", Republish: false}, {Price: "1274.93000", Volume: "2.43100000", Time: "1669902403.994400", Republish: true}}, Bids: []Object(nil), CheckSum: "2819576441", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.04000", Volume: "0.00000000", Time: "1669902404.224301", Republish: false}, {Price: "1274.97000", Volume: "8.28177245", Time: "1669902403.905939", Republish: true}}, Bids: []Object(nil), CheckSum: "3587672765", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.07000", Volume: "0.00000000", Time: "1669902404.229984", Republish: false}, {Price: "1274.98000", Volume: "3.86114224", Time: "1669902403.667804", Republish: true}}, Bids: []Object(nil), CheckSum: "4164727279", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.76000", Volume: "0.00000000", Time: "1669902404.232807", Republish: false}, {Price: "1272.40000", Volume: "4.04378535", Time: "1669902403.905724", Republish: true}}, CheckSum: "1619880483", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.79000", Volume: "0.00000000", Time: "1669902404.240225", Republish: false}, {Price: "1272.39000", Volume: "12.18179960", Time: "1669902403.521893", Republish: true}}, CheckSum: "3995764177", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.84000", Volume: "0.15698956", Time: "1669902404.240733", Republish: false}}, CheckSum: "3865383426", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.71000", Volume: "7.80118122", Time: "1669902404.259689", Republish: false}}, CheckSum: "3726506271", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.13000", Volume: "0.00000000", Time: "1669902404.270642", Republish: false}, {Price: "1274.99000", Volume: "5.51607274", Time: "1669902403.588453", Republish: true}}, Bids: []Object(nil), CheckSum: "3142683171", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.40000", Volume: "0.00000000", Time: "1669902404.289905", Republish: false}, {Price: "1272.39000", Volume: "12.18179960", Time: "1669902403.521893", Republish: true}}, CheckSum: "3641551806", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.92000", Volume: "1.24425015", Time: "1669902404.323209", Republish: false}}, Bids: []Object(nil), CheckSum: "2784583446", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.92000", Volume: "0.00000000", Time: "1669902404.339851", Republish: false}, {Price: "1274.99000", Volume: "5.51607274", Time: "1669902403.588453", Republish: true}}, Bids: []Object(nil), CheckSum: "3641551806", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.93000", Volume: "0.00000000", Time: "1669902404.357052", Republish: false}, {Price: "1275.00000", Volume: "61.40249992", Time: "1669902401.643491", Republish: true}}, Bids: []Object(nil), CheckSum: "3362675611", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.85000", Volume: "2.50000000", Time: "1669902404.456422", Republish: false}}, CheckSum: "1558110789", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.56000", Volume: "23.55579802", Time: "1669902404.456473", Republish: false}}, CheckSum: "2660934152", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.85000", Volume: "2.57209226", Time: "1669902404.460095", Republish: false}}, CheckSum: "1735447414", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.86000", Volume: "2.00000000", Time: "1669902404.482505", Republish: false}}, CheckSum: "834929713", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.06000", Volume: "0.00000000", Time: "1669902404.482569", Republish: false}, {Price: "1275.04000", Volume: "1.25000000", Time: "1669902400.112111", Republish: true}, {Price: "1274.21000", Volume: "3.72191238", Time: "1669902404.482585", Republish: false}}, Bids: []Object(nil), CheckSum: "2763346392", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.85000", Volume: "7.57209226", Time: "1669902404.482672", Republish: false}}, CheckSum: "2186206796", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.86000", Volume: "2.15698956", Time: "1669902404.482901", Republish: false}}, CheckSum: "3460653625", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.57000", Volume: "0.00000000", Time: "1669902404.482915", Republish: false}, {Price: "1272.49000", Volume: "5.00000000", Time: "1669902402.874149", Republish: true}}, CheckSum: "581278027", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.84000", Volume: "0.00000000", Time: "1669902404.484646", Republish: false}, {Price: "1272.39000", Volume: "12.18179960", Time: "1669902403.521893", Republish: true}}, CheckSum: "4189599930", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.83000", Volume: "0.00000000", Time: "1669902404.507485", Republish: false}, {Price: "1272.24000", Volume: "26.96094786", Time: "1669902404.208796", Republish: true}}, CheckSum: "2480919997", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.91000", Volume: "5.50000000", Time: "1669902404.508447", Republish: false}}, CheckSum: "2012579915", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.03000", Volume: "0.00000000", Time: "1669902404.512382", Republish: false}, {Price: "1275.04000", Volume: "1.25000000", Time: "1669902400.112111", Republish: true}}, Bids: []Object(nil), CheckSum: "3689542351", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.86000", Volume: "2.00000000", Time: "1669902404.514003", Republish: false}}, CheckSum: "798466735", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.20000", Volume: "0.15598391", Time: "1669902404.514397", Republish: false}}, Bids: []Object(nil), CheckSum: "3046624971", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.87000", Volume: "0.15698956", Time: "1669902404.525167", Republish: false}}, CheckSum: "263584879", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.72000", Volume: "4.04378535", Time: "1669902404.529866", Republish: false}}, CheckSum: "2301474306", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.85000", Volume: "2.57209226", Time: "1669902404.534338", Republish: false}}, CheckSum: "2463612725", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.87000", Volume: "0.00000000", Time: "1669902404.535383", Republish: false}, {Price: "1272.49000", Volume: "5.00000000", Time: "1669902402.874149", Republish: true}}, CheckSum: "804366754", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.92000", Volume: "0.15698956", Time: "1669902404.539154", Republish: false}}, CheckSum: "1150868092", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.71000", Volume: "0.03918122", Time: "1669902404.548782", Republish: false}, {Price: "1272.92000", Volume: "7.91898956", Time: "1669902404.548799", Republish: false}}, CheckSum: "2066201279", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.92000", Volume: "7.76200000", Time: "1669902404.572629", Republish: false}}, CheckSum: "2205361623", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.93000", Volume: "0.15698956", Time: "1669902404.573537", Republish: false}}, CheckSum: "3207030720", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.71000", Volume: "0.00000000", Time: "1669902404.713430", Republish: false}, {Price: "1272.56000", Volume: "23.55579802", Time: "1669902404.456473", Republish: true}}, CheckSum: "2151907652", IsSnapshot: false},
		{Asks: []Object{{Price: "1275.00000", Volume: "0.00000000", Time: "1669902404.757742", Republish: false}, {Price: "1275.04000", Volume: "1.25000000", Time: "1669902400.112111", Republish: true}}, Bids: []Object(nil), CheckSum: "613983604", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.94000", Volume: "1.37724786", Time: "1669902404.787943", Republish: false}}, CheckSum: "3743812343", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.95000", Volume: "0.15698956", Time: "1669902404.811781", Republish: false}}, CheckSum: "779463548", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.93000", Volume: "0.00000000", Time: "1669902404.812674", Republish: false}, {Price: "1272.65000", Volume: "58.93178495", Time: "1669902403.509657", Republish: true}}, CheckSum: "2637554395", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.78000", Volume: "1.39917553", Time: "1669902404.860986", Republish: false}}, Bids: []Object(nil), CheckSum: "1289329212", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.78000", Volume: "0.00000000", Time: "1669902404.902408", Republish: false}, {Price: "1275.04000", Volume: "1.25000000", Time: "1669902400.112111", Republish: true}}, Bids: []Object(nil), CheckSum: "2637554395", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.96000", Volume: "0.48084730", Time: "1669902404.905488", Republish: false}}, CheckSum: "2651454106", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.93000", Volume: "6.22500000", Time: "1669902404.969870", Republish: false}}, CheckSum: "375355249", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.78000", Volume: "0.00000000", Time: "1669902404.996731", Republish: false}, {Price: "1272.72000", Volume: "4.04378535", Time: "1669902404.529866", Republish: true}}, CheckSum: "656661256", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.97000", Volume: "1.00000000", Time: "1669902404.997196", Republish: false}}, CheckSum: "2386732874", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.79000", Volume: "1.03805897", Time: "1669902405.017688", Republish: false}}, Bids: []Object(nil), CheckSum: "3434660953", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.98000", Volume: "0.15683091", Time: "1669902405.021980", Republish: false}}, CheckSum: "2712575889", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.95000", Volume: "0.00000000", Time: "1669902405.022039", Republish: false}, {Price: "1272.77000", Volume: "2.98013575", Time: "1669902404.048436", Republish: true}}, CheckSum: "1196716312", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.20000", Volume: "4.20115319", Time: "1669902405.044788", Republish: false}}, Bids: []Object(nil), CheckSum: "1111393989", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.96000", Volume: "2.77480339", Time: "1669902405.046011", Republish: false}}, CheckSum: "3084271799", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.20000", Volume: "4.04516928", Time: "1669902405.069780", Republish: false}}, Bids: []Object(nil), CheckSum: "1824964892", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.19000", Volume: "0.15598391", Time: "1669902405.069820", Republish: false}}, Bids: []Object(nil), CheckSum: "4087100690", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.92000", Volume: "0.00000000", Time: "1669902405.148858", Republish: false}, {Price: "1272.72000", Volume: "4.04378535", Time: "1669902404.529866", Republish: true}, {Price: "1272.97000", Volume: "8.76200000", Time: "1669902405.148889", Republish: false}}, CheckSum: "1944326551", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.98000", Volume: "2.41288227", Time: "1669902405.292365", Republish: false}}, CheckSum: "40602593", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.90000", Volume: "0.00000000", Time: "1669902405.292959", Republish: false}, {Price: "1274.99000", Volume: "5.51607274", Time: "1669902403.588453", Republish: true}, {Price: "1274.52000", Volume: "3.15221383", Time: "1669902405.292980", Republish: false}}, Bids: []Object(nil), CheckSum: "3320395706", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.99000", Volume: "12.17605794", Time: "1669902405.295216", Republish: false}}, CheckSum: "1142493595", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.99000", Volume: "12.33288885", Time: "1669902405.317215", Republish: false}}, CheckSum: "129015225", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.98000", Volume: "2.25605136", Time: "1669902405.317685", Republish: false}}, CheckSum: "816404681", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.99000", Volume: "12.17605794", Time: "1669902405.318994", Republish: false}}, CheckSum: "1929434347", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.99000", Volume: "13.17605794", Time: "1669902405.319868", Republish: false}}, CheckSum: "3871816827", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.80000", Volume: "23.29000000", Time: "1669902405.335476", Republish: false}}, Bids: []Object(nil), CheckSum: "4066547813", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.19000", Volume: "0.00000000", Time: "1669902405.342070", Republish: false}, {Price: "1274.99000", Volume: "5.51607274", Time: "1669902403.588453", Republish: true}}, Bids: []Object(nil), CheckSum: "3109850767", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.97000", Volume: "7.76200000", Time: "1669902405.360710", Republish: false}}, CheckSum: "3568817674", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.80000", Volume: "0.00000000", Time: "1669902405.384402", Republish: false}, {Price: "1275.04000", Volume: "1.25000000", Time: "1669902400.112111", Republish: true}}, Bids: []Object(nil), CheckSum: "1030710202", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.85000", Volume: "3.12413933", Time: "1669902405.394006", Republish: false}}, CheckSum: "2652724925", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.20000", Volume: "16.20966464", Time: "1669902405.405018", Republish: false}}, Bids: []Object(nil), CheckSum: "1595310008", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.77000", Volume: "0.48013575", Time: "1669902405.444946", Republish: false}}, CheckSum: "3742031443", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.00000", Volume: "2.50000000", Time: "1669902405.445260", Republish: false}}, CheckSum: "836675112", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.85000", Volume: "2.57209226", Time: "1669902405.446015", Republish: false}}, CheckSum: "293049296", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.11000", Volume: "2.50000000", Time: "1669902405.446349", Republish: false}}, Bids: []Object(nil), CheckSum: "3726847353", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.00000", Volume: "2.57951150", Time: "1669902405.470295", Republish: false}}, CheckSum: "3681192903", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.00000", Volume: "7.57951150", Time: "1669902405.471143", Republish: false}}, CheckSum: "1601929561", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.01000", Volume: "2.00000000", Time: "1669902405.471778", Republish: false}}, CheckSum: "2888325765", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.86000", Volume: "0.00000000", Time: "1669902405.472183", Republish: false}, {Price: "1272.85000", Volume: "2.57209226", Time: "1669902405.446015", Republish: true}}, CheckSum: "1958018041", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.21000", Volume: "0.00000000", Time: "1669902405.529438", Republish: false}, {Price: "1275.04000", Volume: "1.25000000", Time: "1669902400.112111", Republish: true}}, Bids: []Object(nil), CheckSum: "534579802", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.99000", Volume: "36.42605794", Time: "1669902405.535037", Republish: false}}, CheckSum: "1030875576", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.20000", Volume: "12.16449536", Time: "1669902405.572968", Republish: false}}, Bids: []Object(nil), CheckSum: "3511263386", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.00000", Volume: "7.50000000", Time: "1669902405.615188", Republish: false}}, CheckSum: "1010203787", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.01000", Volume: "4.26915968", Time: "1669902405.684112", Republish: false}}, CheckSum: "454144109", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.20000", Volume: "0.00000000", Time: "1669902405.699634", Republish: false}, {Price: "1275.30000", Volume: "1.78523696", Time: "1669902404.530309", Republish: true}}, Bids: []Object(nil), CheckSum: "2944213073", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.02000", Volume: "5.00000000", Time: "1669902405.709535", Republish: false}}, CheckSum: "2282756051", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.02000", Volume: "5.07628095", Time: "1669902405.713684", Republish: false}}, CheckSum: "1478584606", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.97000", Volume: "0.00000000", Time: "1669902405.750486", Republish: false}, {Price: "1272.85000", Volume: "2.57209226", Time: "1669902405.446015", Republish: true}, {Price: "1273.03000", Volume: "7.76200000", Time: "1669902405.750510", Republish: false}}, CheckSum: "2841469680", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.02000", Volume: "0.07628095", Time: "1669902405.777442", Republish: false}}, CheckSum: "2341954265", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.52000", Volume: "1.78523696", Time: "1669902405.788331", Republish: false}, {Price: "1274.96000", Volume: "1.36697687", Time: "1669902405.788353", Republish: false}}, Bids: []Object(nil), CheckSum: "2659105042", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.79000", Volume: "0.00000000", Time: "1669902405.883585", Republish: false}, {Price: "1275.30000", Volume: "1.78523696", Time: "1669902404.530309", Republish: true}}, Bids: []Object(nil), CheckSum: "3097656935", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.95000", Volume: "1.91650646", Time: "1669902405.904628", Republish: false}}, Bids: []Object(nil), CheckSum: "108920218", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.11000", Volume: "0.00000000", Time: "1669902405.936339", Republish: false}, {Price: "1275.30000", Volume: "1.78523696", Time: "1669902404.530309", Republish: true}}, Bids: []Object(nil), CheckSum: "2466214385", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.99000", Volume: "0.00000000", Time: "1669902405.956369", Republish: false}, {Price: "1275.69000", Volume: "0.08890337", Time: "1669902404.530163", Republish: true}}, Bids: []Object(nil), CheckSum: "3761101596", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.90000", Volume: "3.72191238", Time: "1669902405.960224", Republish: false}}, Bids: []Object(nil), CheckSum: "2531677535", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.04000", Volume: "1.00000000", Time: "1669902405.961869", Republish: false}}, CheckSum: "4165791359", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.04000", Volume: "7.22500000", Time: "1669902405.962085", Republish: false}}, CheckSum: "4141850149", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.90000", Volume: "7.71270111", Time: "1669902405.983614", Republish: false}}, Bids: []Object(nil), CheckSum: "2570896953", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.22000", Volume: "0.00000000", Time: "1669902406.000279", Republish: false}, {Price: "1275.69000", Volume: "0.08890337", Time: "1669902404.530163", Republish: true}}, Bids: []Object(nil), CheckSum: "3513420640", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.93000", Volume: "0.00000000", Time: "1669902406.009405", Republish: false}, {Price: "1272.91000", Volume: "5.50000000", Time: "1669902404.508447", Republish: true}}, CheckSum: "1160547181", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.99000", Volume: "35.42605794", Time: "1669902406.009997", Republish: false}}, CheckSum: "4084125432", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.98000", Volume: "0.00000000", Time: "1669902406.047049", Republish: false}, {Price: "1275.96000", Volume: "78.37270831", Time: "1669902398.285332", Republish: true}}, Bids: []Object(nil), CheckSum: "3468220961", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.05000", Volume: "4.16630783", Time: "1669902406.047102", Republish: false}}, CheckSum: "1531839482", IsSnapshot: false},
		{Asks: []Object{{Price: "1275.95000", Volume: "5.75381173", Time: "1669902406.047660", Republish: false}}, Bids: []Object(nil), CheckSum: "3213446116", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.52000", Volume: "0.00000000", Time: "1669902406.048274", Republish: false}, {Price: "1275.96000", Volume: "78.37270831", Time: "1669902398.285332", Republish: true}}, Bids: []Object(nil), CheckSum: "2352283827", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.91000", Volume: "0.00000000", Time: "1669902406.048500", Republish: false}, {Price: "1276.00000", Volume: "1.25000000", Time: "1669902383.215558", Republish: true}}, Bids: []Object(nil), CheckSum: "2747940740", IsSnapshot: false},
		{Asks: []Object{{Price: "1275.69000", Volume: "0.00000000", Time: "1669902406.049970", Republish: false}, {Price: "1276.36000", Volume: "0.21749417", Time: "1669902404.011270", Republish: true}}, Bids: []Object(nil), CheckSum: "2505105620", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.05000", Volume: "9.66630783", Time: "1669902406.086791", Republish: false}}, CheckSum: "1746616005", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.89000", Volume: "12.15791166", Time: "1669902406.125218", Republish: false}}, Bids: []Object(nil), CheckSum: "3824194371", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.06000", Volume: "0.14662260", Time: "1669902406.128250", Republish: false}}, CheckSum: "3245434816", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.02000", Volume: "0.00000000", Time: "1669902406.211196", Republish: false}, {Price: "1272.94000", Volume: "1.37724786", Time: "1669902404.787943", Republish: true}}, CheckSum: "1103833741", IsSnapshot: false},
		{Asks: []Object{{Price: "1275.57000", Volume: "3.68734191", Time: "1669902406.310837", Republish: false}}, Bids: []Object(nil), CheckSum: "788075958", IsSnapshot: false},
		{Asks: []Object{{Price: "1275.00000", Volume: "3.72168949", Time: "1669902406.324368", Republish: false}}, Bids: []Object(nil), CheckSum: "2645805010", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.03000", Volume: "0.00000000", Time: "1669902406.352702", Republish: false}, {Price: "1272.91000", Volume: "5.50000000", Time: "1669902404.508447", Republish: true}, {Price: "1273.06000", Volume: "7.90862260", Time: "1669902406.352723", Republish: false}}, CheckSum: "3664794344", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.88000", Volume: "0.48084730", Time: "1669902406.353511", Republish: false}}, Bids: []Object(nil), CheckSum: "3417984912", IsSnapshot: false},
		{Asks: []Object{{Price: "1275.00000", Volume: "0.00000000", Time: "1669902406.370487", Republish: false}, {Price: "1275.95000", Volume: "5.75381173", Time: "1669902406.047660", Republish: true}}, Bids: []Object(nil), CheckSum: "3150946599", IsSnapshot: false},
		{Asks: []Object{{Price: "1275.11000", Volume: "23.29000000", Time: "1669902406.371246", Republish: false}}, Bids: []Object(nil), CheckSum: "1911788631", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.07000", Volume: "2.50000000", Time: "1669902406.435166", Republish: false}}, CheckSum: "3138928474", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.11000", Volume: "2.50000000", Time: "1669902406.435279", Republish: false}}, Bids: []Object(nil), CheckSum: "1369393517", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1272.99000", Volume: "12.17605794", Time: "1669902406.453111", Republish: false}, {Price: "1273.08000", Volume: "23.25000000", Time: "1669902406.453126", Republish: false}}, CheckSum: "3429481298", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.08000", Volume: "25.25000000", Time: "1669902406.461396", Republish: false}}, CheckSum: "2192001657", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.87000", Volume: "2.00000000", Time: "1669902406.461762", Republish: false}}, Bids: []Object(nil), CheckSum: "2502755534", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.01000", Volume: "2.26915968", Time: "1669902406.462321", Republish: false}}, CheckSum: "1341802899", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.07000", Volume: "7.50000000", Time: "1669902406.463292", Republish: false}}, CheckSum: "1761723911", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.47000", Volume: "12.16191828", Time: "1669902406.473535", Republish: false}}, Bids: []Object(nil), CheckSum: "2200635864", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.05000", Volume: "5.50000000", Time: "1669902406.482051", Republish: false}}, CheckSum: "95232443", IsSnapshot: false},
		{Asks: []Object(nil), Bids: []Object{{Price: "1273.07000", Volume: "2.50000000", Time: "1669902406.499508", Republish: false}}, CheckSum: "592883247", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.89000", Volume: "0.00000000", Time: "1669902406.503260", Republish: false}, {Price: "1275.11000", Volume: "23.29000000", Time: "1669902406.371246", Republish: true}}, Bids: []Object(nil), CheckSum: "1172665353", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.87000", Volume: "25.54402527", Time: "1669902406.583127", Republish: false}}, Bids: []Object(nil), CheckSum: "3039926281", IsSnapshot: false},
		{Asks: []Object{{Price: "1275.11000", Volume: "0.00000000", Time: "1669902406.622406", Republish: false}, {Price: "1275.30000", Volume: "1.78523696", Time: "1669902404.530309", Republish: true}}, Bids: []Object(nil), CheckSum: "2789250059", IsSnapshot: false},
		{Asks: []Object{{Price: "1274.11000", Volume: "0.00000000", Time: "1669902406.708357", Republish: false}, {Price: "1275.57000", Volume: "3.68734191", Time: "1669902406.310837", Republish: true}}, Bids: []Object(nil), CheckSum: "3246434297", IsSnapshot: false},
	}
}
