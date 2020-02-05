package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

// MetaField is a the metadata for a dataset field
type MetaField struct {
	Description string  `json:"description" yaml:"description" firestore:"description"`
	Dtype       string  `json:"dtype" yaml:"dtype" firestore:"dtype"`
	Required    bool    `json:"required" yaml:"required" firestore:"required"`
	FillRate    float64 `json:"fill_rate" yaml:"fill_rate" firestore:"fill_rate"`
}

// MetaFile is the expected structure of the YAML metadata documents
type MetaFile struct {
	Name               string                 `json:"name" yaml:"name" firestore:"name"`
	Version            string                 `json:"version" yaml:"version" firestore:"version"`
	Description        string                 `json:"description" yaml:"description" firestore:"description"`
	Class              string                 `json:"class" yaml:"class" firestore:"class"`
	Format             string                 `json:"format" yaml:"format" firestore:"format"`
	Provenance         Provenance             `json:"provenance" yaml:"provenance" firestore:"provenance"`
	Labels             []string               `json:"labels" yaml:"labels" firestore:"labels"`
	Fields             yaml.MapSlice        `json:"fields" yaml:"fields" firestore:"fields"`
	AdditionalMetadata map[string]interface{} `json:"additional_metadata" yaml:"additional_metadata" firestore:"additional_metadata"`
}

// Provenance is a sub struct of the expected YAML metadata describing the provenance of the data
type Provenance struct {
	Source    string `json:"source" yaml:"source" firestore:"source"`
	ValidFrom string `json:"valid_from" yaml:"valid_from" firestore:"valid_from"`
	ValidTo   string `json:"valid_to" yaml:"valid_to" firestore:"valid_to"`
}

func main() {
	fmt.Println("Hello, playground")
	metaData := `
name: Main Company Information
version: '3'
description: Contains key information about the companies, including name, date created, premises type and industry classification.
class: raw
format: text
provenance:
  source: Creditsafe
  valid_from:
  valid_to:
labels: [main, name, created, established, premises, industry, website, employees, turnover]
fields:
  identifier:
    description: Identifier for table. In this case NLCI01.
    dtype: STRING
    required: true
    fill_rate: 1
  creditsafe_global_ID:
    description: Creditsafe global ID.
    dtype: STRING
    required: true
    fill_rate: 1
  company_trading_number:
    description: Unique internal ID.
    dtype: STRING
    required: true
    fill_rate: 1
  company_name:
    description: Name of the company.
    dtype: STRING
    fill_rate: 0.999997
  alphaname:
    description: This looks like a code with no spaces that represents the company name.
    dtype: STRING
    fill_rate: 0.999991
  alphaname_short:
    description: A shorter version of alphaname.
    dtype: STRING
    fill_rate: 0.922565
  established_year_month:
    description: The month and year in which the company was set up at address (YYYYMM).
    dtype: FLOAT
    fill_rate: 0.495307
  linked_limited_company_registration_number:
    description: Linked limited company registration number.
    dtype: STRING
    fill_rate: 0.015632
  group_ID:
    description: Can be used to group companies together (Asda stores for example).
    dtype: FLOAT
    fill_rate: 0.153799
  premises_type:
    description: Refers to lookup table NLLU01, which looks up the type of business premises.
    dtype: STRING
    fill_rate: 0.997950
  SIC_code_2007:
    description: Refers to lookup table NLLU02, which looks up 2007 SIC code descriptions.
    dtype: FLOAT
    fill_rate: 0.819356
  website_address:
    description: Website address.
    dtype: STRING
    fill_rate: 0.512487
  employee_number_estimate:
    description: Estimates the number of employees.
    dtype: FLOAT
    fill_rate: 0.841518
  turnover_band:
    description: Turnover band.
    dtype: STRING
    fill_rate: 0.153980

`
	var meta MetaFile
	_ = yaml.Unmarshal([]byte(metaData), &meta)
	for _, x := range meta.Fields.Content {
	    var field interface{}
	    fmt.Println(x.Decode(field))
	}
}
