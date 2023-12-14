package smd

import (
	"context"
	"database/sql"
	"smile-service/entities"
)

const (
	execSPDataChecklistFreezer = `exec [sp_smile_smd_freezer_insert]$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20`
)

func InsertDataChecklistFreezer(ctx context.Context, db *sql.DB,
	CustomerNo string,
	BillTo string,
	ShipTo string,
	FrezerAvailable int,
	CreatedBy string,
	CreatedByIP string,
	SerialNo string,
	Merk int,
	NoteMerk string,
	Capacity int,
	Location int,
	Access int,
	NoteAccess string,
	Power int,
	NotePower string,
	Condition int,
	NoteCondition string,
	OutsideCondition int,
	NoteOutsideCondition string,
	Suhu int) error {
	rows, err := db.QueryContext(ctx, execSPDataChecklistFreezer,
		CustomerNo,
		BillTo,
		ShipTo,
		FrezerAvailable,
		CreatedBy,
		CreatedByIP,
		SerialNo,
		Merk,
		NoteMerk,
		Capacity,
		Location,
		Access,
		NoteAccess,
		Power,
		NotePower,
		Condition,
		NoteCondition,
		OutsideCondition,
		NoteOutsideCondition,
		Suhu,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func InsertChecklistFreezer(ctx context.Context, db *sql.DB, checklistFreezer *entities.DataSMDFreezerInsert) error {

	err := InsertDataChecklistFreezer(ctx, db,
		checklistFreezer.Header.CustomerNo, checklistFreezer.Header.BillTo, checklistFreezer.Header.ShipTo, checklistFreezer.Header.FrezerAvailable,
		checklistFreezer.Header.CreatedBy, checklistFreezer.Header.CreatedByIP, checklistFreezer.DetailFreezer.SerialNo,
		checklistFreezer.DetailFreezer.Merk, checklistFreezer.DetailFreezer.NoteMerk, checklistFreezer.DetailFreezer.Capacity,
		checklistFreezer.DetailFreezer.Location, checklistFreezer.DetailFreezer.Access, checklistFreezer.DetailFreezer.NoteAccess,
		checklistFreezer.DetailFreezer.Power, checklistFreezer.DetailFreezer.NotePower, checklistFreezer.DetailFreezer.Condition,
		checklistFreezer.DetailFreezer.NoteCondition, checklistFreezer.DetailFreezer.OutsideCondition, checklistFreezer.DetailFreezer.NoteOutsideCondition,
		checklistFreezer.DetailFreezer.Suhu,
	)

	if err != nil {
		return err
	}

	return nil
}
