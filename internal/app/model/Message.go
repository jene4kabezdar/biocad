package model

import (
	"strconv"

	"github.com/jene4kabezdar/biocad/internal/app/store"
)

const pageSize int = 5

type Message struct {
	Number    int    `json:"Number"`
	Mqtt      string `json:"Mqtt"`
	Invid     string `json:"Invid"`
	Unit_guid string `json:"Unit_guid"`
	Msg_id    string `json:"Msg_id"`
	Text      string `json:"Text"`
	Context   string `json:"Context"`
	Class     string `json:"Class"`
	Level     int    `json:"Level"`
	Area      string `json:"Area"`
	Addr      string `json:"Addr"`
	Block     string `json:"Block"`
	Type      string `json:"Type"`
	Bit       string `json:"Bit"`
	InvertBit string `json:"InvertBit"`
}

func (m *Message) Add(store store.Store) (*Message, error) {
	var number int
	if err := store.DB.QueryRow(
		`INSERT INTO messages 
			(mqtt, invid, unit_guid, msg_id, text, context, class, level, area, 
				addr, block, type, bit, invert_bit)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
			RETURNING n`,
		m.Mqtt, m.Invid, m.Unit_guid, m.Msg_id, m.Text, m.Context, m.Class, m.Level,
		m.Area, m.Addr, m.Block, m.Type, m.Bit, m.InvertBit,
	).Scan(&m.Number); err != nil {
		return nil, err
	}

	m.Number = number

	return m, nil
}

func GetMessages(store store.Store, number int) ([]Message, error) {
	rows, err := store.DB.Query(`SELECT n, mqtt, invid, unit_guid, msg_id, text, 
									context, class, level, area, addr, block, 
									type, bit, invert_bit FROM messages LIMIT $1 OFFSET $2`,
									pageSize, number*pageSize)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []Message

	for rows.Next() {
		var mes Message
		if err := rows.Scan(&mes.Number, &mes.Mqtt, &mes.Invid, &mes.Unit_guid,
			&mes.Msg_id, &mes.Text, &mes.Context, &mes.Class, &mes.Level, &mes.Area,
			&mes.Addr, &mes.Block, &mes.Type, &mes.Bit, &mes.InvertBit); err != nil {
			return messages, err
		}
		messages = append(messages, mes)
	}

	if err = rows.Err(); err != nil {
		return messages, err
	}

	return messages, nil
}

func CreateMessagesByRows(rows [][]string) ([]Message, error) {
	res := make([]Message, len(rows)-2)

	for i := 2; i < len(rows); i++ {
		level, err := strconv.Atoi(rows[i][8])
		if err != nil {
			return nil, err
		}

		message := Message{
			Mqtt:      rows[i][1],
			Invid:     rows[i][2],
			Unit_guid: rows[i][3],
			Msg_id:    rows[i][4],
			Text:      rows[i][5],
			Context:   rows[i][6],
			Class:     rows[i][7],
			Level:     level,
			Area:      rows[i][9],
			Addr:      rows[i][10],
			Block:     rows[i][11],
			Type:      rows[i][12],
			Bit:       rows[i][13],
			InvertBit: rows[i][14],
		}
		res[i-2] = message
	}
	return res, nil
}
