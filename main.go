package main

import (
	"fmt"
	"strings"
	"time"
)

type Ticket struct {
    Ticket string
    User   string
    Status string
    Date   time.Time
}

func getTickets(lines []string) []Ticket{
	layout := "2006-01-02"

	tickets := make([]Ticket, 0, len(lines))

	for _, line := range lines{
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, "TICKET-") {
			continue
		}

		ticket_field := strings.Split(line, "_")

		if (len(ticket_field) != 4){
			continue
		}

		timeStamp, err := time.Parse(layout, ticket_field[3])

		if (err != nil){
			continue
		}

		switch ticket_field[2] {
		case "Готово":
			break
		case "В работе":
			break
		case "Не будет сделано":
			break
		default:
			continue
		}

		ticket := Ticket{
			Ticket: ticket_field[0],
			User: ticket_field[1],
			Status: ticket_field[2],
			Date: timeStamp,
		}

		tickets = append(tickets, ticket)
	}

	return tickets
}

func filterUser(tickets []Ticket, user *string) []Ticket{
	if (user == nil){
		return tickets
	}

	filter_tickets := make([]Ticket, 0, len(tickets))

	for i := range tickets{
		if (tickets[i].User == *user){
			filter_tickets = append(filter_tickets, tickets[i])
		}
	}

	return filter_tickets
}

func filterStatus(tickets []Ticket, status *string) []Ticket{
	if (status == nil){
		return tickets
	}

	filter_tickets := make([]Ticket, 0, len(tickets))

	for i := range tickets{
		if (tickets[i].Status == *status){
			filter_tickets = append(filter_tickets, tickets[i])
		}
	}

	return filter_tickets
}


func GetTasks(text string, user *string, status *string) []Ticket{
	lines := strings.Split(text, "\n")

	tickets := getTickets(lines)

	tickets = filterStatus(filterUser(tickets, user), status)

	return tickets
}

func main(){
	text := `
TICKET-12345_Паша Попов_Готово_2024-01-01
TICKET-12346_Иван Иванов_В работе_2024-01-02
TICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03
TICKET-12348_Паша Попов_В работе_2024-01-04`

	user := "Паша Попов"
	tickets := GetTasks(text, &user, nil)

	for _, line := range tickets{
		fmt.Println(line.Ticket, line.User, line.Status, line.Date)
	}
}
