// Где работаем!
package main
 
// Импорт библеотек
import ("fmt"
		"net/http")

func home_page(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Ia rabotay ohyet")
}

// Отслежевание переходов по страницам
func HandleRec()  {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8080", nil)
}

// Основная функция где собираем дерево функций
func main()  {
	HandleRec()
}