# RFC 6238 TOTP Generator

This project implements a Time-Based One-Time Password (TOTP) generator in Go, strictly following RFC 6238, with custom parameters required by the challenge specification.

The generated TOTP is used as the password for HTTP Basic Authentication.

## TOTP Configuration (Challenge-Specific)
| Parameter      | Value                          |
| -------------- | ------------------------------ |
| RFC            | RFC 6238                       |
| Time Step (X)  | 30 seconds                     |
| T0             | 0                              |
| Hash Algorithm | HMAC-SHA-512                   |
| OTP Length     | 10 digits                      |
| Secret         | `email + "HENNGECHALLENGE004"` |


## Algorithm Summary
Current Time
   ↓
Time Counter = floor((T - T0) / 30)
   ↓
HMAC-SHA-512(secret, counter)
   ↓
Dynamic Truncation (RFC 6238)
   ↓
Modulo 10¹⁰
   ↓
10-digit TOTP
