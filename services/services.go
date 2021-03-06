/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package services

import "net/http"

type Route struct {
	Name    string           `json:"Name"`
	Method  string           `json:"Method"`
	Pattern string           `json:"Pattern"`
	Handler http.HandlerFunc `json:"Handler"`
}

type RouteCollection []Route
