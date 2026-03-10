package main 

import ("fmt"
	"time"
	"math"
)

func main(){

	var date time.Time
	var str string 

	fmt.Scanln(&str)
	date, _ = time.Parse("02.01.2006", str)

	var str1, str2, str3 string 
	fmt.Scan(&str1, &str2, &str3)

	var a, b, c float64 
	fmt.Scan(&a, &b, &c)

	new_date := date.AddDate(0, 0, 15)

	var sum float64 
	sum = a + b + c

	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия" + 
	" сформировала документы по факту выполненной вами работы.\n" +
	"Дата подписания договора: %s." +
	" Просим вас подойти в офис в любое удобное для вас время в этот день.\n" + 
	"Общая сумма выплат составит %d руб. %d коп.\n\n" + 
	"С уважением,\nГл. бух. Иванов А.Е.\n", str2, str1, str3, new_date.Format("02.01.2006"), int(math.Floor(sum)), 
	int(math.Round((sum - math.Floor(sum)) * 100)))

}