// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VersionVersionMatrix VersionMatrix represents set of possible product versions.
//
// swagger:model versionVersionMatrix
type VersionVersionMatrix struct {

	// backup
	Backup map[string]VersionVersion `json:"backup,omitempty"`

	// haproxy
	Haproxy map[string]VersionVersion `json:"haproxy,omitempty"`

	// log collector
	LogCollector map[string]VersionVersion `json:"logCollector,omitempty"`

	// mongod
	Mongod map[string]VersionVersion `json:"mongod,omitempty"`

	// operator
	Operator map[string]VersionVersion `json:"operator,omitempty"`

	// pgbackrest
	Pgbackrest map[string]VersionVersion `json:"pgbackrest,omitempty"`

	// pgbackrest repo
	PgbackrestRepo map[string]VersionVersion `json:"pgbackrestRepo,omitempty"`

	// pgbadger
	Pgbadger map[string]VersionVersion `json:"pgbadger,omitempty"`

	// pgbouncer
	Pgbouncer map[string]VersionVersion `json:"pgbouncer,omitempty"`

	// pmm
	Pmm map[string]VersionVersion `json:"pmm,omitempty"`

	// postgresql
	Postgresql map[string]VersionVersion `json:"postgresql,omitempty"`

	// proxysql
	Proxysql map[string]VersionVersion `json:"proxysql,omitempty"`

	// psmdb operator
	PsmdbOperator map[string]VersionVersion `json:"psmdbOperator,omitempty"`

	// pxc
	Pxc map[string]VersionVersion `json:"pxc,omitempty"`

	// pxc operator
	PxcOperator map[string]VersionVersion `json:"pxcOperator,omitempty"`
}

// Validate validates this version version matrix
func (m *VersionVersionMatrix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHaproxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogCollector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMongod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePgbackrest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePgbackrestRepo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePgbadger(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePgbouncer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostgresql(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePsmdbOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePxc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePxcOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionVersionMatrix) validateBackup(formats strfmt.Registry) error {
	if swag.IsZero(m.Backup) { // not required
		return nil
	}

	for k := range m.Backup {

		if err := validate.Required("backup"+"."+k, "body", m.Backup[k]); err != nil {
			return err
		}
		if val, ok := m.Backup[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validateHaproxy(formats strfmt.Registry) error {
	if swag.IsZero(m.Haproxy) { // not required
		return nil
	}

	for k := range m.Haproxy {

		if err := validate.Required("haproxy"+"."+k, "body", m.Haproxy[k]); err != nil {
			return err
		}
		if val, ok := m.Haproxy[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validateLogCollector(formats strfmt.Registry) error {
	if swag.IsZero(m.LogCollector) { // not required
		return nil
	}

	for k := range m.LogCollector {

		if err := validate.Required("logCollector"+"."+k, "body", m.LogCollector[k]); err != nil {
			return err
		}
		if val, ok := m.LogCollector[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validateMongod(formats strfmt.Registry) error {
	if swag.IsZero(m.Mongod) { // not required
		return nil
	}

	for k := range m.Mongod {

		if err := validate.Required("mongod"+"."+k, "body", m.Mongod[k]); err != nil {
			return err
		}
		if val, ok := m.Mongod[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	for k := range m.Operator {

		if err := validate.Required("operator"+"."+k, "body", m.Operator[k]); err != nil {
			return err
		}
		if val, ok := m.Operator[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validatePgbackrest(formats strfmt.Registry) error {
	if swag.IsZero(m.Pgbackrest) { // not required
		return nil
	}

	for k := range m.Pgbackrest {

		if err := validate.Required("pgbackrest"+"."+k, "body", m.Pgbackrest[k]); err != nil {
			return err
		}
		if val, ok := m.Pgbackrest[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validatePgbackrestRepo(formats strfmt.Registry) error {
	if swag.IsZero(m.PgbackrestRepo) { // not required
		return nil
	}

	for k := range m.PgbackrestRepo {

		if err := validate.Required("pgbackrestRepo"+"."+k, "body", m.PgbackrestRepo[k]); err != nil {
			return err
		}
		if val, ok := m.PgbackrestRepo[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validatePgbadger(formats strfmt.Registry) error {
	if swag.IsZero(m.Pgbadger) { // not required
		return nil
	}

	for k := range m.Pgbadger {

		if err := validate.Required("pgbadger"+"."+k, "body", m.Pgbadger[k]); err != nil {
			return err
		}
		if val, ok := m.Pgbadger[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validatePgbouncer(formats strfmt.Registry) error {
	if swag.IsZero(m.Pgbouncer) { // not required
		return nil
	}

	for k := range m.Pgbouncer {

		if err := validate.Required("pgbouncer"+"."+k, "body", m.Pgbouncer[k]); err != nil {
			return err
		}
		if val, ok := m.Pgbouncer[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validatePmm(formats strfmt.Registry) error {
	if swag.IsZero(m.Pmm) { // not required
		return nil
	}

	for k := range m.Pmm {

		if err := validate.Required("pmm"+"."+k, "body", m.Pmm[k]); err != nil {
			return err
		}
		if val, ok := m.Pmm[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validatePostgresql(formats strfmt.Registry) error {
	if swag.IsZero(m.Postgresql) { // not required
		return nil
	}

	for k := range m.Postgresql {

		if err := validate.Required("postgresql"+"."+k, "body", m.Postgresql[k]); err != nil {
			return err
		}
		if val, ok := m.Postgresql[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validateProxysql(formats strfmt.Registry) error {
	if swag.IsZero(m.Proxysql) { // not required
		return nil
	}

	for k := range m.Proxysql {

		if err := validate.Required("proxysql"+"."+k, "body", m.Proxysql[k]); err != nil {
			return err
		}
		if val, ok := m.Proxysql[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validatePsmdbOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.PsmdbOperator) { // not required
		return nil
	}

	for k := range m.PsmdbOperator {

		if err := validate.Required("psmdbOperator"+"."+k, "body", m.PsmdbOperator[k]); err != nil {
			return err
		}
		if val, ok := m.PsmdbOperator[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validatePxc(formats strfmt.Registry) error {
	if swag.IsZero(m.Pxc) { // not required
		return nil
	}

	for k := range m.Pxc {

		if err := validate.Required("pxc"+"."+k, "body", m.Pxc[k]); err != nil {
			return err
		}
		if val, ok := m.Pxc[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) validatePxcOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.PxcOperator) { // not required
		return nil
	}

	for k := range m.PxcOperator {

		if err := validate.Required("pxcOperator"+"."+k, "body", m.PxcOperator[k]); err != nil {
			return err
		}
		if val, ok := m.PxcOperator[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this version version matrix based on the context it is used
func (m *VersionVersionMatrix) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHaproxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogCollector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMongod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePgbackrest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePgbackrestRepo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePgbadger(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePgbouncer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePmm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostgresql(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxysql(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePsmdbOperator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePxc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePxcOperator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionVersionMatrix) contextValidateBackup(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Backup {

		if val, ok := m.Backup[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidateHaproxy(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Haproxy {

		if val, ok := m.Haproxy[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidateLogCollector(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.LogCollector {

		if val, ok := m.LogCollector[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidateMongod(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Mongod {

		if val, ok := m.Mongod[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidateOperator(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Operator {

		if val, ok := m.Operator[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidatePgbackrest(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Pgbackrest {

		if val, ok := m.Pgbackrest[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidatePgbackrestRepo(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.PgbackrestRepo {

		if val, ok := m.PgbackrestRepo[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidatePgbadger(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Pgbadger {

		if val, ok := m.Pgbadger[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidatePgbouncer(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Pgbouncer {

		if val, ok := m.Pgbouncer[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidatePmm(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Pmm {

		if val, ok := m.Pmm[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidatePostgresql(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Postgresql {

		if val, ok := m.Postgresql[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidateProxysql(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Proxysql {

		if val, ok := m.Proxysql[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidatePsmdbOperator(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.PsmdbOperator {

		if val, ok := m.PsmdbOperator[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidatePxc(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Pxc {

		if val, ok := m.Pxc[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *VersionVersionMatrix) contextValidatePxcOperator(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.PxcOperator {

		if val, ok := m.PxcOperator[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionVersionMatrix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionVersionMatrix) UnmarshalBinary(b []byte) error {
	var res VersionVersionMatrix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
