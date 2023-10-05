// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

package extstatefulset

const (
	StatefulSetTargetType = "com.steadybit.extension_kubernetes.kubernetes-statefulset"
	statefulSetIcon       = "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9Im5vIj8+CjwhLS0gQ3JlYXRlZCB3aXRoIElua3NjYXBlIChodHRwOi8vd3d3Lmlua3NjYXBlLm9yZy8pIC0tPgoKPHN2ZwogICB4bWxuczpkYz0iaHR0cDovL3B1cmwub3JnL2RjL2VsZW1lbnRzLzEuMS8iCiAgIHhtbG5zOmNjPSJodHRwOi8vY3JlYXRpdmVjb21tb25zLm9yZy9ucyMiCiAgIHhtbG5zOnJkZj0iaHR0cDovL3d3dy53My5vcmcvMTk5OS8wMi8yMi1yZGYtc3ludGF4LW5zIyIKICAgeG1sbnM6c3ZnPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIKICAgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIgogICB4bWxuczpzb2RpcG9kaT0iaHR0cDovL3NvZGlwb2RpLnNvdXJjZWZvcmdlLm5ldC9EVEQvc29kaXBvZGktMC5kdGQiCiAgIHhtbG5zOmlua3NjYXBlPSJodHRwOi8vd3d3Lmlua3NjYXBlLm9yZy9uYW1lc3BhY2VzL2lua3NjYXBlIgogICB3aWR0aD0iMTguMDM1MzM0bW0iCiAgIGhlaWdodD0iMTcuNTAwMzc4bW0iCiAgIHZpZXdCb3g9IjAgMCAxOC4wMzUzMzQgMTcuNTAwMzc4IgogICB2ZXJzaW9uPSIxLjEiCiAgIGlkPSJzdmcxMzgyNiIKICAgaW5rc2NhcGU6dmVyc2lvbj0iMC45MSByMTM3MjUiCiAgIHNvZGlwb2RpOmRvY25hbWU9InN0cy5zdmciPgogIDxkZWZzCiAgICAgaWQ9ImRlZnMxMzgyMCIgLz4KICA8c29kaXBvZGk6bmFtZWR2aWV3CiAgICAgaWQ9ImJhc2UiCiAgICAgcGFnZWNvbG9yPSIjZmZmZmZmIgogICAgIGJvcmRlcmNvbG9yPSIjNjY2NjY2IgogICAgIGJvcmRlcm9wYWNpdHk9IjEuMCIKICAgICBpbmtzY2FwZTpwYWdlb3BhY2l0eT0iMC4wIgogICAgIGlua3NjYXBlOnBhZ2VzaGFkb3c9IjIiCiAgICAgaW5rc2NhcGU6em9vbT0iOCIKICAgICBpbmtzY2FwZTpjeD0iLTIuMDkwMDA0IgogICAgIGlua3NjYXBlOmN5PSIyOC43NTIyMzkiCiAgICAgaW5rc2NhcGU6ZG9jdW1lbnQtdW5pdHM9Im1tIgogICAgIGlua3NjYXBlOmN1cnJlbnQtbGF5ZXI9ImxheWVyMSIKICAgICBzaG93Z3JpZD0iZmFsc2UiCiAgICAgaW5rc2NhcGU6d2luZG93LXdpZHRoPSIxNDQwIgogICAgIGlua3NjYXBlOndpbmRvdy1oZWlnaHQ9Ijc3NSIKICAgICBpbmtzY2FwZTp3aW5kb3cteD0iMCIKICAgICBpbmtzY2FwZTp3aW5kb3cteT0iMSIKICAgICBpbmtzY2FwZTp3aW5kb3ctbWF4aW1pemVkPSIxIgogICAgIGZpdC1tYXJnaW4tdG9wPSIwIgogICAgIGZpdC1tYXJnaW4tbGVmdD0iMCIKICAgICBmaXQtbWFyZ2luLXJpZ2h0PSIwIgogICAgIGZpdC1tYXJnaW4tYm90dG9tPSIwIiAvPgogIDxtZXRhZGF0YQogICAgIGlkPSJtZXRhZGF0YTEzODIzIj4KICAgIDxyZGY6UkRGPgogICAgICA8Y2M6V29yawogICAgICAgICByZGY6YWJvdXQ9IiI+CiAgICAgICAgPGRjOmZvcm1hdD5pbWFnZS9zdmcreG1sPC9kYzpmb3JtYXQ+CiAgICAgICAgPGRjOnR5cGUKICAgICAgICAgICByZGY6cmVzb3VyY2U9Imh0dHA6Ly9wdXJsLm9yZy9kYy9kY21pdHlwZS9TdGlsbEltYWdlIiAvPgogICAgICAgIDxkYzp0aXRsZSAvPgogICAgICA8L2NjOldvcms+CiAgICA8L3JkZjpSREY+CiAgPC9tZXRhZGF0YT4KICA8ZwogICAgIGlua3NjYXBlOmxhYmVsPSJDYWxxdWUgMSIKICAgICBpbmtzY2FwZTpncm91cG1vZGU9ImxheWVyIgogICAgIGlkPSJsYXllcjEiCiAgICAgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTAuOTkyNjI2MzgsLTEuMTc0MTgxKSI+CiAgICA8ZwogICAgICAgaWQ9Imc3MCIKICAgICAgIHRyYW5zZm9ybT0ibWF0cml4KDEuMDE0ODg4NywwLDAsMS4wMTQ4ODg3LDE2LjkzNzQyNCwtMi42OTg3MjYpIj4KICAgICAgPHBhdGgKICAgICAgICAgaW5rc2NhcGU6ZXhwb3J0LXlkcGk9IjI1MC41NSIKICAgICAgICAgaW5rc2NhcGU6ZXhwb3J0LXhkcGk9IjI1MC41NSIKICAgICAgICAgaW5rc2NhcGU6ZXhwb3J0LWZpbGVuYW1lPSJuZXcucG5nIgogICAgICAgICBpbmtzY2FwZTpjb25uZWN0b3ItY3VydmF0dXJlPSIwIgogICAgICAgICBpZD0icGF0aDMwNTUiCiAgICAgICAgIGQ9Im0gLTYuODQ5MjAxNSw0LjI3MjQ2NjggYSAxLjExOTEyNTUsMS4xMDk5NjcxIDAgMCAwIC0wLjQyODg4MTgsMC4xMDg1MzAzIGwgLTUuODUyNDAzNywyLjc5NjMzOTQgYSAxLjExOTEyNTUsMS4xMDk5NjcxIDAgMCAwIC0wLjYwNTUyNCwwLjc1Mjk3NTkgbCAtMS40NDM4MjgsNi4yODEyODQ2IGEgMS4xMTkxMjU1LDEuMTA5OTY3MSAwIDAgMCAwLjE1MTk0MywwLjg1MTAyOCAxLjExOTEyNTUsMS4xMDk5NjcxIDAgMCAwIDAuMDYzNjIsMC4wODgzMiBsIDQuMDUwOCw1LjAzNjU1NSBhIDEuMTE5MTI1NSwxLjEwOTk2NzEgMCAwIDAgMC44NzQ5NzksMC40MTc2NTQgbCA2LjQ5NjEwMTEsLTAuMDAxNSBhIDEuMTE5MTI1NSwxLjEwOTk2NzEgMCAwIDAgMC44NzQ5Nzg4LC0wLjQxNjkwNiBMIDEuMzgxODg3MiwxNS4xNDk0NTMgQSAxLjExOTEyNTUsMS4xMDk5NjcxIDAgMCAwIDEuNTk4MTk4NiwxNC4yMTAxMDQgTCAwLjE1MjEyNjU3LDcuOTI4ODE1NCBBIDEuMTE5MTI1NSwxLjEwOTk2NzEgMCAwIDAgLTAuNDUzMzk3OTQsNy4xNzU4Mzk2IEwgLTYuMzA2NTQ5Niw0LjM4MDk5NzEgQSAxLjExOTEyNTUsMS4xMDk5NjcxIDAgMCAwIC02Ljg0OTIwMTUsNC4yNzI0NjY4IFoiCiAgICAgICAgIHN0eWxlPSJmaWxsOiMzMjZjZTU7ZmlsbC1vcGFjaXR5OjE7c3Ryb2tlOm5vbmU7c3Ryb2tlLXdpZHRoOjA7c3Ryb2tlLW1pdGVybGltaXQ6NDtzdHJva2UtZGFzaGFycmF5Om5vbmU7c3Ryb2tlLW9wYWNpdHk6MSIgLz4KICAgICAgPHBhdGgKICAgICAgICAgaWQ9InBhdGgzMDU0LTItOSIKICAgICAgICAgZD0iTSAtNi44NTIzNDM1LDMuODE3NjM3MiBBIDEuMTgxNDMwNCwxLjE3MTc2MiAwIDAgMCAtNy4zMDQ0Mjg0LDMuOTMyOTA0IGwgLTYuMTc4NzQyNiwyLjk1MTI3NTggYSAxLjE4MTQzMDQsMS4xNzE3NjIgMCAwIDAgLTAuNjM5MjA2LDAuNzk0ODkxIGwgLTEuNTIzOTE1LDYuNjMwODI4MiBhIDEuMTgxNDMwNCwxLjE3MTc2MiAwIDAgMCAwLjE2MDE3NSwwLjg5ODkzIDEuMTgxNDMwNCwxLjE3MTc2MiAwIDAgMCAwLjA2NzM2LDAuMDkyODEgbCA0LjI3NjA5NCw1LjMxNzIzNiBhIDEuMTgxNDMwNCwxLjE3MTc2MiAwIDAgMCAwLjkyMzYzLDAuNDQwODU4IGwgNi44NTc2MTg4LC0wLjAwMTUgYSAxLjE4MTQzMDQsMS4xNzE3NjIgMCAwIDAgMC45MjM2MzA4LC0wLjQ0MDExIGwgNC4yNzQ1OTY2LC01LjMxNzk4NSBhIDEuMTgxNDMwNCwxLjE3MTc2MiAwIDAgMCAwLjIyODI4OCwtMC45OTA5OTMgTCAwLjUzODk0NDM5LDcuNjc3NTczOCBBIDEuMTgxNDMwNCwxLjE3MTc2MiAwIDAgMCAtMC4xMDAyNjEwMSw2Ljg4MzQzMTMgTCAtNi4yNzkwMDM3LDMuOTMyMTU1NSBBIDEuMTgxNDMwNCwxLjE3MTc2MiAwIDAgMCAtNi44NTIzNDM1LDMuODE3NjM3MiBaIG0gMC4wMDI5OSwwLjQ1NTA3ODkgYSAxLjExOTEyNTUsMS4xMDk5NjcxIDAgMCAxIDAuNTQyNjUxNywwLjEwODUzMDMgbCA1Ljg1MzE1MTY5LDIuNzk0ODQyNSBBIDEuMTE5MTI1NSwxLjEwOTk2NzEgMCAwIDEgMC4xNTE5NzgxMSw3LjkyOTA2NDggTCAxLjU5ODA1MSwxNC4yMTAzNSBhIDEuMTE5MTI1NSwxLjEwOTk2NzEgMCAwIDEgLTAuMjE2MzEyMywwLjkzOTM0OCBsIC00LjA0OTMwMzIsNS4wMzczMDQgYSAxLjExOTEyNTUsMS4xMDk5NjcxIDAgMCAxIC0wLjg3NDk3ODksMC40MTY5MDYgbCAtNi40OTYxMDA2LDAuMDAxNSBhIDEuMTE5MTI1NSwxLjEwOTk2NzEgMCAwIDEgLTAuODc0OTc5LC0wLjQxNzY1MiBsIC00LjA1MDgsLTUuMDM2NTU0IGEgMS4xMTkxMjU1LDEuMTA5OTY3MSAwIDAgMSAtMC4wNjM2MiwtMC4wODgzMiAxLjExOTEyNTUsMS4xMDk5NjcxIDAgMCAxIC0wLjE1MTk0MiwtMC44NTEwMjggbCAxLjQ0MzgyNywtNi4yODEyODUzIGEgMS4xMTkxMjU1LDEuMTA5OTY3MSAwIDAgMSAwLjYwNTUyNCwtMC43NTI5NzU4IGwgNS44NTI0MDM2LC0yLjc5NjMzOTUgYSAxLjExOTEyNTUsMS4xMDk5NjcxIDAgMCAxIDAuNDI4ODgxOSwtMC4xMDg1MzAzIHoiCiAgICAgICAgIHN0eWxlPSJjb2xvcjojMDAwMDAwO2ZvbnQtc3R5bGU6bm9ybWFsO2ZvbnQtdmFyaWFudDpub3JtYWw7Zm9udC13ZWlnaHQ6bm9ybWFsO2ZvbnQtc3RyZXRjaDpub3JtYWw7Zm9udC1zaXplOm1lZGl1bTtsaW5lLWhlaWdodDpub3JtYWw7Zm9udC1mYW1pbHk6U2FuczstaW5rc2NhcGUtZm9udC1zcGVjaWZpY2F0aW9uOlNhbnM7dGV4dC1pbmRlbnQ6MDt0ZXh0LWFsaWduOnN0YXJ0O3RleHQtZGVjb3JhdGlvbjpub25lO3RleHQtZGVjb3JhdGlvbi1saW5lOm5vbmU7bGV0dGVyLXNwYWNpbmc6bm9ybWFsO3dvcmQtc3BhY2luZzpub3JtYWw7dGV4dC10cmFuc2Zvcm06bm9uZTtkaXJlY3Rpb246bHRyO3dyaXRpbmctbW9kZTpsci10YjtiYXNlbGluZS1zaGlmdDpiYXNlbGluZTt0ZXh0LWFuY2hvcjpzdGFydDtkaXNwbGF5OmlubGluZTtvdmVyZmxvdzp2aXNpYmxlO3Zpc2liaWxpdHk6dmlzaWJsZTtmaWxsOiNmZmZmZmY7ZmlsbC1vcGFjaXR5OjE7ZmlsbC1ydWxlOm5vbnplcm87c3Ryb2tlOm5vbmU7c3Ryb2tlLXdpZHRoOjA7c3Ryb2tlLW1pdGVybGltaXQ6NDtzdHJva2UtZGFzaGFycmF5Om5vbmU7bWFya2VyOm5vbmU7ZW5hYmxlLWJhY2tncm91bmQ6YWNjdW11bGF0ZSIKICAgICAgICAgaW5rc2NhcGU6Y29ubmVjdG9yLWN1cnZhdHVyZT0iMCIgLz4KICAgIDwvZz4KICAgIDxnCiAgICAgICBpZD0iZzQyMjEiCiAgICAgICB0cmFuc2Zvcm09InRyYW5zbGF0ZSgwLjI2ODc3NzYyLDEuMDkyMzEwNSkiPgogICAgICA8cGF0aAogICAgICAgICBzdHlsZT0iZmlsbDpub25lO2ZpbGwtcnVsZTpldmVub2RkO3N0cm9rZTojZmZmZmZmO3N0cm9rZS13aWR0aDowLjUyOTE0NTg0O3N0cm9rZS1saW5lY2FwOnNxdWFyZTtzdHJva2UtbGluZWpvaW46cm91bmQ7c3Ryb2tlLW1pdGVybGltaXQ6MTA7c3Ryb2tlLWRhc2hhcnJheToxLjU4NzQzNzYxLCAxLjU4NzQzNzYxO3N0cm9rZS1kYXNob2Zmc2V0OjMuNjY2OTgwNzQ7c3Ryb2tlLW9wYWNpdHk6MSIKICAgICAgICAgaW5rc2NhcGU6Y29ubmVjdG9yLWN1cnZhdHVyZT0iMCIKICAgICAgICAgZD0ibSA4LjA1MzAzMzMsNS4xMjkwNzU2IDYuNTI1MDA2NywwIDAsNC41ODMzMzUgLTYuNTI1MDA2NywwIHoiCiAgICAgICAgIGlkPSJwYXRoMTIzNCIgLz4KICAgICAgPGcKICAgICAgICAgaWQ9Imc0MjEzIj4KICAgICAgICA8cGF0aAogICAgICAgICAgIGlkPSJwYXRoMTIzOCIKICAgICAgICAgICBkPSJtIDYuNTE0Mjg0OSw2LjU0MDM4NzYgNi41MjUwMDcxLDAgMCw0LjU4MzMzNTQgLTYuNTI1MDA3MSwwIHoiCiAgICAgICAgICAgaW5rc2NhcGU6Y29ubmVjdG9yLWN1cnZhdHVyZT0iMCIKICAgICAgICAgICBzdHlsZT0iZmlsbDojMzI2Y2U1O2ZpbGwtb3BhY2l0eToxO2ZpbGwtcnVsZTpldmVub2RkO3N0cm9rZTojZmZmZmZmO3N0cm9rZS13aWR0aDowLjUyOTE0NTg0O3N0cm9rZS1saW5lY2FwOnNxdWFyZTtzdHJva2UtbGluZWpvaW46cm91bmQ7c3Ryb2tlLW1pdGVybGltaXQ6MTA7c3Ryb2tlLWRhc2hhcnJheToxLjU4NzQzNzYxLCAxLjU4NzQzNzYxO3N0cm9rZS1kYXNob2Zmc2V0OjMuODc4NjM4OTg7c3Ryb2tlLW9wYWNpdHk6MSIgLz4KICAgICAgICA8cGF0aAogICAgICAgICAgIGlkPSJwYXRoMTI0MCIKICAgICAgICAgICBkPSJtIDQuOTc1NTU3OCw3Ljk1MTY5ODQgNi41MjQ5OTEyLDAgMCw0LjU4MzMzNDYgLTYuNTI0OTkxMiwwIHoiCiAgICAgICAgICAgaW5rc2NhcGU6Y29ubmVjdG9yLWN1cnZhdHVyZT0iMCIKICAgICAgICAgICBzdHlsZT0iZmlsbDojZmZmZmZmO2ZpbGwtcnVsZTpldmVub2RkO3N0cm9rZTpub25lO3N0cm9rZS13aWR0aDowLjI2NDU4MzMyO3N0cm9rZS1saW5lY2FwOnNxdWFyZTtzdHJva2UtbWl0ZXJsaW1pdDoxMCIgLz4KICAgICAgICA8cGF0aAogICAgICAgICAgIGlkPSJwYXRoMTI0MiIKICAgICAgICAgICBkPSJtIDQuOTc1NTU3OCw3Ljk1MTY5ODQgNi41MjQ5OTEyLDAgMCw0LjU4MzMzNDYgLTYuNTI0OTkxMiwwIHoiCiAgICAgICAgICAgaW5rc2NhcGU6Y29ubmVjdG9yLWN1cnZhdHVyZT0iMCIKICAgICAgICAgICBzdHlsZT0iZmlsbDpub25lO2ZpbGwtcnVsZTpldmVub2RkO3N0cm9rZTojZmZmZmZmO3N0cm9rZS13aWR0aDowLjUyOTE2Njc7c3Ryb2tlLWxpbmVjYXA6YnV0dDtzdHJva2UtbGluZWpvaW46cm91bmQ7c3Ryb2tlLW1pdGVybGltaXQ6MTA7c3Ryb2tlLW9wYWNpdHk6MSIgLz4KICAgICAgICA8cGF0aAogICAgICAgICAgIGlkPSJwYXRoMTI0NiIKICAgICAgICAgICBkPSJtIDUuNTA4NzE5OCw5LjIyNjA0ODIgMCwwIGMgMCwtMC40Mjk0MDU0IDEuMjIxODg4MSwtMC43Nzc1MDU0IDIuNzI5MTYxNSwtMC43Nzc1MDU0IDEuNTA3Mjc0MywwIDIuNzI5MTc4NywwLjM0ODEgMi43MjkxNzg3LDAuNzc3NTA1NCBsIDAsMCBjIDAsMC40Mjk0MDU0IC0xLjIyMTkwNDQsMC43Nzc1MDY4IC0yLjcyOTE3ODcsMC43Nzc1MDY4IC0xLjUwNzI3MzQsMCAtMi43MjkxNjE1LC0wLjM0ODEwMTQgLTIuNzI5MTYxNSwtMC43Nzc1MDY4IHoiCiAgICAgICAgICAgaW5rc2NhcGU6Y29ubmVjdG9yLWN1cnZhdHVyZT0iMCIKICAgICAgICAgICBzdHlsZT0iZmlsbDojZmZmZmZmO2ZpbGwtb3BhY2l0eToxO2ZpbGwtcnVsZTpldmVub2RkO3N0cm9rZTpub25lO3N0cm9rZS13aWR0aDowLjI2NDU4MzMyO3N0cm9rZS1saW5lY2FwOnNxdWFyZTtzdHJva2UtbWl0ZXJsaW1pdDoxMCIgLz4KICAgICAgICA8cGF0aAogICAgICAgICAgIGlkPSJwYXRoMTI0OCIKICAgICAgICAgICBkPSJtIDEwLjk2NzA2Nyw5LjIyNjA0ODIgMCwwIGMgMCwwLjQyOTQwNTQgLTEuMjIxOTA2NCwwLjc3NzUwNjggLTIuNzI5MTc3NywwLjc3NzUwNjggLTEuNTA3MjczNCwwIC0yLjcyOTE2MTQsLTAuMzQ4MTAxNCAtMi43MjkxNjE0LC0wLjc3NzUwNjggbCAwLDAgYyAwLC0wLjQyOTQwNTQgMS4yMjE4ODgsLTAuNzc3NTA1NCAyLjcyOTE2MTQsLTAuNzc3NTA1NCAxLjUwNzI3MzMsMCAyLjcyOTE3NzcsMC4zNDgxIDIuNzI5MTc3NywwLjc3NzUwNTQgbCAwLDIuMTAzMzIyOCBjIDAsMC40Mjk0MDUgLTEuMjIxOTA2NCwwLjc3NzUwNiAtMi43MjkxNzc3LDAuNzc3NTA2IC0xLjUwNzI3MzQsMCAtMi43MjkxNjE0LC0wLjM0ODEwMSAtMi43MjkxNjE0LC0wLjc3NzUwNiBsIDAsLTIuMTAzMzIyOCIKICAgICAgICAgICBpbmtzY2FwZTpjb25uZWN0b3ItY3VydmF0dXJlPSIwIgogICAgICAgICAgIHN0eWxlPSJmaWxsOiMwMDAwMDA7ZmlsbC1vcGFjaXR5OjA7ZmlsbC1ydWxlOmV2ZW5vZGQ7c3Ryb2tlOm5vbmU7c3Ryb2tlLXdpZHRoOjAuMjY0NTgzMzI7c3Ryb2tlLWxpbmVjYXA6c3F1YXJlO3N0cm9rZS1taXRlcmxpbWl0OjEwIiAvPgogICAgICAgIDxwYXRoCiAgICAgICAgICAgaWQ9InBhdGgxMjUwIgogICAgICAgICAgIGQ9Im0gMTAuOTY3MDY3LDkuMjI2MDQ4MiAwLDAgYyAwLDAuNDI5NDA1NCAtMS4yMjE5MDY0LDAuNzc3NTA2OCAtMi43MjkxNzc3LDAuNzc3NTA2OCAtMS41MDcyNzM0LDAgLTIuNzI5MTYxNCwtMC4zNDgxMDE0IC0yLjcyOTE2MTQsLTAuNzc3NTA2OCBsIDAsMCBjIDAsLTAuNDI5NDA1NCAxLjIyMTg4OCwtMC43Nzc1MDU0IDIuNzI5MTYxNCwtMC43Nzc1MDU0IDEuNTA3MjczMywwIDIuNzI5MTc3NywwLjM0ODEgMi43MjkxNzc3LDAuNzc3NTA1NCBsIDAsMi4xMDMzMjI4IGMgMCwwLjQyOTQwNSAtMS4yMjE5MDY0LDAuNzc3NTA2IC0yLjcyOTE3NzcsMC43Nzc1MDYgLTEuNTA3MjczNCwwIC0yLjcyOTE2MTQsLTAuMzQ4MTAxIC0yLjcyOTE2MTQsLTAuNzc3NTA2IGwgMCwtMi4xMDMzMjI4IgogICAgICAgICAgIGlua3NjYXBlOmNvbm5lY3Rvci1jdXJ2YXR1cmU9IjAiCiAgICAgICAgICAgc3R5bGU9ImZpbGw6bm9uZTtmaWxsLXJ1bGU6ZXZlbm9kZDtzdHJva2U6IzMyNmNlNTtzdHJva2Utd2lkdGg6MC41MjkxNjY3O3N0cm9rZS1saW5lY2FwOmJ1dHQ7c3Ryb2tlLWxpbmVqb2luOnJvdW5kO3N0cm9rZS1taXRlcmxpbWl0OjEwO3N0cm9rZS1vcGFjaXR5OjEiIC8+CiAgICAgIDwvZz4KICAgIDwvZz4KICA8L2c+Cjwvc3ZnPgo="
)
