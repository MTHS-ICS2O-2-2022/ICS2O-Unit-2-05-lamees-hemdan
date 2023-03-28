// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Lamees Hemdan
// Created on: March 2023
// This file contains the JS functions for index.html

'use strict'
/**
 * This function calculates your income tax
 */
  function myButtonClicked() {
    // input
    const hours = parseFloat(document.getElementById('hours').value)
    const salary = parseFloat(document.getElementById('salary').value)
    const TAX_RATE = 0.18

    // process
    const pay = (hours * salary) * (1 - TAX_RATE)
    const governmentPay = (hours * salary) * TAX_RATE
    
    // output
    document.getElementById('pay').innerHTML = `Your pay will be: $ ${pay.toFixed(2)}`
    document.getElementById('tax').innerHTML = `Government will Take: $ ${governmentPay.toFixed(2)}`
  }
