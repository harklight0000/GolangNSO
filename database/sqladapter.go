package database

import (
	"fmt"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"
	"gorm.io/gorm"
	. "nso/ainterfaces"
	"nso/errs"
	"nso/logging"
	"nso/sqlplugins"
	"reflect"
	"strconv"
	"strings"
)

type sqlDatabaseAdapter struct {
	*gorm.DB
}

func (d *sqlDatabaseAdapter) FindOneOfTable(table string, result interface{}, condition ICondition) error {
	if condition != nil {
		if err := d.DB.Table(table).Where(condition.ToSQLCondition()).First(result).Error; err != nil {
			return eris.Wrap(err, "FindOne error")
		}
	} else {
		return eris.Wrap(eris.New("Must pass condition to that function"), "FindOne error")
	}
	return nil
}

func NewSQLDatabaseAdapter(db *gorm.DB) IDatabase {
	return &sqlDatabaseAdapter{db}
}

func (d *sqlDatabaseAdapter) Unwrap(v interface{}) {
	if reflect.ValueOf(v).Kind() == reflect.Ptr {
		reflect.ValueOf(v).Elem().Set(reflect.ValueOf(d.DB))
	} else {
		logging.Logger.Info("Unsupported type", zap.String("type", reflect.ValueOf(v).Type().String()))
	}
}

func (d *sqlDatabaseAdapter) Create(value interface{}) error {
	table := getTableName(value)
	createResult := d.DB.Table(table).Create(value)
	if createResult.Error != nil {
		return eris.Wrap(createResult.Error, "Create error")
	}
	return nil
}

func (d *sqlDatabaseAdapter) Update(table string, value interface{}, condition ICondition) error {
	if err := d.DB.Table(table).Where(condition.ToSQLCondition()).UpdateColumns(value).Error; err != nil {
		return eris.Wrap(err, "Update error")
	}
	return nil
}

func (d *sqlDatabaseAdapter) Delete(table string, condition ICondition) error {
	if err := d.DB.Table(table).Raw("DELETE FROM " + table + " WHERE " + condition.ToSQLCondition()).Error; err != nil {
		return eris.Wrap(err, "Delete error")
	}
	return nil
}

func (d *sqlDatabaseAdapter) FindOne(result interface{}, condition ICondition) error {
	table := getTableName(result)
	if condition != nil {
		if err := d.DB.Table(table).Where(condition.ToSQLCondition()).First(result).Error; err != nil {
			return eris.Wrap(err, "FindOne error")
		}
	} else {
		return eris.Wrap(eris.New("Must pass condition to that function"), "FindOne error")
	}
	return nil
}

func (d *sqlDatabaseAdapter) FindManyWithConditionAndOption(result interface{}, condition ICondition, option IOption) error {

	table := getTableName(result)
	var cond string
	if condition == nil {
		cond = ""
	} else {
		cond = strings.Trim(condition.ToSQLCondition(), " ")
	}
	tx := d.DB.Table(table)

	if cond != "" {
		tx.Where(cond)
	} else if option != nil {
		switch option.(type) {
		case *sqlplugins.LimitOption:
			limit := option.(*sqlplugins.LimitOption)
			count, err := strconv.Atoi(fmt.Sprintf("%v", limit.Value))
			if err != nil {
				logging.Logger.Info("Error converting limit to int!! " + errs.ToString(err))
				count = 1
			}
			tx.Limit(count)
		case *sqlplugins.AscOption:
			tx.Order(option.(*sqlplugins.AscOption).Name)
		case *sqlplugins.DescOption:
			tx.Order(option.(*sqlplugins.DescOption).Name + " desc")
		default:
			logging.Logger.Info("Unsupported option", zap.String("type", reflect.TypeOf(option).String()))
		}
	}

	if err := tx.Find(result).Error; err != nil {
		return eris.Wrap(err, "FindManyWithConditionAndOption error of table "+table)
	}
	return nil
}

func (d *sqlDatabaseAdapter) FindManyWithOption(result interface{}, options IOption) error {
	return d.FindManyWithConditionAndOption(result, sqlplugins.DefaultCondition(), options)
}

func (d *sqlDatabaseAdapter) FindMany(result interface{}, condition ICondition) error {
	return d.FindManyWithConditionAndOption(result, condition, nil)
}

func (d *sqlDatabaseAdapter) FindAll(result interface{}) error {
	return d.FindManyWithConditionAndOption(result, nil, nil)
}

func (d *sqlDatabaseAdapter) Close() error {
	logging.Logger.Info("Closed database connection")
	return nil
}
