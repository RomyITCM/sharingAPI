package smd

import (
	"context"
	"database/sql"
	"smile-service/entities"
)

const (
	execSPDataChecklistFreezerUpdate = `exec [sp_smile_smd_freezer_update]$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17`
)

func UpdateDataChecklistFreezer(ctx context.Context, db *sql.DB,
	TransNo string,
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
	Suhu int,
	CreatedBy string,
	CreatedByIP string) error {
	rows, err := db.QueryContext(ctx, execSPDataChecklistFreezerUpdate,
		TransNo,
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
		CreatedBy,
		CreatedByIP,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func UpdateChecklistFreezer(ctx context.Context, db *sql.DB, checklistFreezer *entities.DataSMDFreezerUpdate) error {

	err := UpdateDataChecklistFreezer(ctx, db,
		checklistFreezer.TransNo.TransNo, checklistFreezer.DetailFreezer.SerialNo,
		checklistFreezer.DetailFreezer.Merk, checklistFreezer.DetailFreezer.NoteMerk, checklistFreezer.DetailFreezer.Capacity,
		checklistFreezer.DetailFreezer.Location, checklistFreezer.DetailFreezer.Access, checklistFreezer.DetailFreezer.NoteAccess,
		checklistFreezer.DetailFreezer.Power, checklistFreezer.DetailFreezer.NotePower, checklistFreezer.DetailFreezer.Condition,
		checklistFreezer.DetailFreezer.NoteCondition, checklistFreezer.DetailFreezer.OutsideCondition, checklistFreezer.DetailFreezer.NoteOutsideCondition,
		checklistFreezer.DetailFreezer.Suhu, checklistFreezer.TransNo.CreatedBy, checklistFreezer.TransNo.CreatedByIP,
	)

	if err != nil {
		return err
	}

	return nil
}
