#!/bin/bash

genny -in=ovirtObjectMethods.template -out=ovirtObjectMethods.go gen "OvirtObjectType=VM,Cluster,DataCenter,Template" -pkg ovirtapi
