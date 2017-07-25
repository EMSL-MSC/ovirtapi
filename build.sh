#!/bin/bash

genny -in=ovirtObjectMethods.template -out=ovirtObjectMethods.go gen "OvirtObjectType=Vm,Cluster,DataCenter" -pkg ovirtapi
