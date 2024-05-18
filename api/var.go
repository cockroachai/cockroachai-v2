package api

var (
	Props = `
  {
    "props": {
        "pageProps": {
            "user": {
                "id": "user-xyhelper",
                "name": "admin@closeai.biz",
                "email": "admin@closeai.biz",
                "image": "https://s.gravatar.com/avatar/558db47f25d89a95df170b4bde9fd72f?s=480\u0026r=pg\u0026d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fli.png",
                "picture": "https://s.gravatar.com/avatar/558db47f25d89a95df170b4bde9fd72f?s=480\u0026r=pg\u0026d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fli.png",
                "idp": "auth0",
                "iat": 1715650700,
                "mfa": false,
                "groups": [],
                "intercom_hash": "30fd0a0ada1c07ce526be7c3d54c22904b80fa7e2713d978630e979e4315cf67"
            },
            "serviceStatus": {},
            "serviceAnnouncement": {
                "public": {},
                "paid": {}
            },
            "userCountry": "US",
            "userRegion": "New York",
            "userRegionCode": "NY",
            "cfConnectingIp": "8.8.8.8",
            "cfIpCity": "Buffalo",
            "cfTimezone": "America/New_York",
            "userLocale": "en",
            "statsig": {
                "payload": {
                    "feature_gates": {
                        "29898433": {
                            "name": "29898433",
                            "value": true,
                            "rule_id": "5hlR2AjP7h2nJXOxZHKzKH:100.00:2",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "48895201": {
                            "name": "48895201",
                            "value": false,
                            "rule_id": "3i3FRD5a2bDACgRX2SCksw",
                            "secondary_exposures": []
                        },
                        "73891701": {
                            "name": "73891701",
                            "value": true,
                            "rule_id": "4M51IcwRuyJAV7JLf2RvzR:100.00:1",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:chatgpt_unpaid_users",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:chatgpt_paid_users",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "81143276": {
                            "name": "81143276",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": []
                        },
                        "87687712": {
                            "name": "87687712",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:european_legal_restricted",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "174366048": {
                            "name": "174366048",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:chatgpt_voice_team",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:friends_and_family_nda",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:chatgpt_paid_users",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:chatgpt_unpaid_users",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "chatgpt-web-voice-message-playback",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "211801712": {
                            "name": "211801712",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:european_legal_restricted",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "226799327": {
                            "name": "226799327",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "437245079": {
                            "name": "437245079",
                            "value": true,
                            "rule_id": "1KKPtXmOjxZamlgM2DzXlg",
                            "secondary_exposures": []
                        },
                        "458009956": {
                            "name": "458009956",
                            "value": true,
                            "rule_id": "6LgEwykI8hHnFAQ5EkKTX2:100.00:1",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:plusqa_testers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:scallion_free_preview_testers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:scallion_paid_testers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "544659310": {
                            "name": "544659310",
                            "value": true,
                            "rule_id": "1Opz1br1XsEM6Uajp6F1XU:100.00:3",
                            "secondary_exposures": []
                        },
                        "607723475": {
                            "name": "607723475",
                            "value": false,
                            "rule_id": "2e0cON2LQDfCjkZQ6yzHg5",
                            "secondary_exposures": []
                        },
                        "716475478": {
                            "name": "716475478",
                            "value": true,
                            "rule_id": "48mToGJmOtQ5aRx9MCjOor",
                            "secondary_exposures": []
                        },
                        "716722001": {
                            "name": "716722001",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": []
                        },
                        "730579107": {
                            "name": "730579107",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:european_legal_restricted",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "740954505": {
                            "name": "740954505",
                            "value": false,
                            "rule_id": "2Ey97LJY7QymPfJfxULJwr:10.00:4",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:anonymous_users",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "internal-employee-only-chatgpt-team",
                                    "gateValue": "false",
                                    "ruleID": "4nM2ehmgoDQIv69B0zohb6"
                                },
                                {
                                    "gate": "segment:enterprise",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:plusqa_testers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:scallion_free_preview_testers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:scallion_paid_testers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "771993664": {
                            "name": "771993664",
                            "value": true,
                            "rule_id": "1dDdxrXy1KadHtYusuGRgR",
                            "secondary_exposures": []
                        },
                        "857491970": {
                            "name": "857491970",
                            "value": true,
                            "rule_id": "2LpyNTIfbYC4R6XS6dfdIL",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "859327441": {
                            "name": "859327441",
                            "value": false,
                            "rule_id": "2kLMbBkcSZ8TkzHgzgbyKV",
                            "secondary_exposures": []
                        },
                        "874794369": {
                            "name": "874794369",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:european_legal_restricted",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "1026575464": {
                            "name": "1026575464",
                            "value": true,
                            "rule_id": "1f4cSeenpcmWQ9eMKKQhF5",
                            "secondary_exposures": []
                        },
                        "1056591003": {
                            "name": "1056591003",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:internal_oai_gptv_testers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:chatgpt_unpaid_users",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "chatgpt-mainline-model-transition-unpaid-sahara-v",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "1070272635": {
                            "name": "1070272635",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:friends_and_family_nda",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:feather_memory_trainers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "1123641887": {
                            "name": "1123641887",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:european_legal_restricted",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "chatgpt_anon_chat_enabled",
                                    "gateValue": "true",
                                    "ruleID": "22nVhoL17eyMvGWgFrDfZe"
                                }
                            ]
                        },
                        "1123642877": {
                            "name": "1123642877",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:european_legal_restricted",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "chatgpt_anon_chat_enabled",
                                    "gateValue": "true",
                                    "ruleID": "22nVhoL17eyMvGWgFrDfZe"
                                }
                            ]
                        },
                        "1162285750": {
                            "name": "1162285750",
                            "value": false,
                            "rule_id": "5SD2EhDYytn75oTAjOCSiL",
                            "secondary_exposures": []
                        },
                        "1198030896": {
                            "name": "1198030896",
                            "value": true,
                            "rule_id": "2aaGr86KHS6GUSs6I9F2cV",
                            "secondary_exposures": []
                        },
                        "1426009137": {
                            "name": "1426009137",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:chatgpt_voice_team",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "1568195116": {
                            "name": "1568195116",
                            "value": false,
                            "rule_id": "KKhehNjRxNHoIWKNkqpU2",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "1654633603": {
                            "name": "1654633603",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:anonymous_users",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "internal-employee-only-chatgpt-team",
                                    "gateValue": "false",
                                    "ruleID": "4nM2ehmgoDQIv69B0zohb6"
                                },
                                {
                                    "gate": "segment:enterprise",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:plusqa_testers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:scallion_free_preview_testers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:scallion_paid_testers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "chatgpt_fruit_juice",
                                    "gateValue": "false",
                                    "ruleID": "2Ey97LJY7QymPfJfxULJwr:10.00:4"
                                }
                            ]
                        },
                        "1703773460": {
                            "name": "1703773460",
                            "value": false,
                            "rule_id": "63qNN95X7P6VSHVff2nL76",
                            "secondary_exposures": []
                        },
                        "1820151945": {
                            "name": "1820151945",
                            "value": false,
                            "rule_id": "6trY8Zt53urtkxuGnQToZr",
                            "secondary_exposures": []
                        },
                        "1825130190": {
                            "name": "1825130190",
                            "value": true,
                            "rule_id": "Nef2uMceNUF9U3ZYwSbpD",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:plusqa_testers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "1828361965": {
                            "name": "1828361965",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": []
                        },
                        "1852364556": {
                            "name": "1852364556",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": []
                        },
                        "1880477506": {
                            "name": "1880477506",
                            "value": true,
                            "rule_id": "7fFsLeJu6pyCq9tRWmX9em:100.00:6",
                            "secondary_exposures": []
                        },
                        "1883608273": {
                            "name": "1883608273",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": []
                        },
                        "1923022511": {
                            "name": "1923022511",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:chatgpt_voice_team",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:friends_and_family_nda",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:chatgpt_paid_users",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:chatgpt_unpaid_users",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "1938289220": {
                            "name": "1938289220",
                            "value": false,
                            "rule_id": "1HDaU0d5mGA5CS4wzvOloC",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:contractors",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:sidekick_windows_contractors",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:plusqa_testers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:friends_and_family_nda",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:enterprise",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:chatgpt_paid_users",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "2000076788": {
                            "name": "2000076788",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:contractors",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "2004295676": {
                            "name": "2004295676",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "chatgpt-sonic-internal",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "chatgpt-search-demo-ui",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "2046333385": {
                            "name": "2046333385",
                            "value": false,
                            "rule_id": "4nM2ehmgoDQIv69B0zohb6",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "2166885617": {
                            "name": "2166885617",
                            "value": false,
                            "rule_id": "1ORnWBMVyZ42eAp03fe5Nr",
                            "secondary_exposures": []
                        },
                        "2173548801": {
                            "name": "2173548801",
                            "value": true,
                            "rule_id": "22nVhoL17eyMvGWgFrDfZe",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:european_legal_restricted",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "2438078654": {
                            "name": "2438078654",
                            "value": true,
                            "rule_id": "VNBdYfNSNNUyDdMpaiaZ9:100.00:3",
                            "secondary_exposures": []
                        },
                        "2449630478": {
                            "name": "2449630478",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": []
                        },
                        "2542814321": {
                            "name": "2542814321",
                            "value": false,
                            "rule_id": "6Usmt3Th0mkUkVQJ0oyeST:0.00:1",
                            "secondary_exposures": []
                        },
                        "2546838099": {
                            "name": "2546838099",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "chatgpt-sonic-internal",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "2562876640": {
                            "name": "2562876640",
                            "value": true,
                            "rule_id": "326czTZeZ0RX0ypR0c5Bb6:100.00:15",
                            "secondary_exposures": []
                        },
                        "2689530298": {
                            "name": "2689530298",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:european_legal_restricted",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "2865702999": {
                            "name": "2865702999",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:european_legal_restricted",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "2865703032": {
                            "name": "2865703032",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:european_legal_restricted",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "2893926039": {
                            "name": "2893926039",
                            "value": false,
                            "rule_id": "1XOa3bw1GRog1lbmWKnaF9",
                            "secondary_exposures": []
                        },
                        "2947776398": {
                            "name": "2947776398",
                            "value": false,
                            "rule_id": "2YzNwxxdROEWn1lPwJbYtr",
                            "secondary_exposures": []
                        },
                        "3011278998": {
                            "name": "3011278998",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": []
                        },
                        "3060586957": {
                            "name": "3060586957",
                            "value": false,
                            "rule_id": "1agh2VKHmwqbsORiGxd3v5",
                            "secondary_exposures": []
                        },
                        "3134329717": {
                            "name": "3134329717",
                            "value": false,
                            "rule_id": "2UWRV6bObClvvj8jLkP7kJ",
                            "secondary_exposures": []
                        },
                        "3254335143": {
                            "name": "3254335143",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": []
                        },
                        "3320547351": {
                            "name": "3320547351",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "3376455464": {
                            "name": "3376455464",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:european_legal_restricted",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "3455796462": {
                            "name": "3455796462",
                            "value": true,
                            "rule_id": "YbNn9kuhmMn3DdnzEDc6G",
                            "secondary_exposures": []
                        },
                        "3501550327": {
                            "name": "3501550327",
                            "value": false,
                            "rule_id": "HEU9zVkuwUfmXHhGcckPQ",
                            "secondary_exposures": []
                        },
                        "3511639124": {
                            "name": "3511639124",
                            "value": true,
                            "rule_id": "1y9t5DBS0nz9mqbQRCrmCP:100.00:5",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:plus_qa_-_gizmos",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:store_dogfooding",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:anonymous_users",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:gizmo_store_holdout_exceptions",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "3643162575": {
                            "name": "3643162575",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": []
                        },
                        "3700615661": {
                            "name": "3700615661",
                            "value": false,
                            "rule_id": "66covmutTYx82FWVUlZAqF",
                            "secondary_exposures": []
                        },
                        "3912844760": {
                            "name": "3912844760",
                            "value": false,
                            "rule_id": "default",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:european_legal_restricted",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "3922145230": {
                            "name": "3922145230",
                            "value": false,
                            "rule_id": "14DZA2LumaPqAdCo52CrUB",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:internal_ips",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "3941367332": {
                            "name": "3941367332",
                            "value": true,
                            "rule_id": "3tesvDaP0kRy0JAXNdSJJn:100.00:3",
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_web_sharing_modal_simplified_targeting",
                                    "gateValue": "true",
                                    "ruleID": "48mToGJmOtQ5aRx9MCjOor"
                                }
                            ]
                        },
                        "4226692983": {
                            "name": "4226692983",
                            "value": true,
                            "rule_id": "6sEu91zwlBGSKOqFiNpGlA:100.00:2",
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:enterprise",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:teams_plan",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        }
                    },
                    "dynamic_configs": {
                        "75342665": {
                            "name": "75342665",
                            "value": {
                                "is_anon_chat_enabled": true,
                                "anon_composer_display_variant": "default",
                                "anon-is-spanish-translation-enabled": true,
                                "should_show_anon_login_header_on_desktop": false,
                                "is_anon_chat_enabled_for_new_users_only": false,
                                "is_try_it_first_on_login_page_enabled": false,
                                "is_no_auth_welcome_modal_enabled": false,
                                "no_auth_soft_rate_limit": 1200,
                                "no_auth_hard_rate_limit": 1200,
                                "should_show_no_auth_signup_banner": true,
                                "is_no_auth_welcome_back_modal_enabled": true
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ],
                            "explicit_parameters": [
                                "is_anon_chat_enabled"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "75342698": {
                            "name": "75342698",
                            "value": {
                                "is_anon_chat_enabled": true,
                                "anon_composer_display_variant": "default",
                                "anon-is-spanish-translation-enabled": true,
                                "should_show_anon_login_header_on_desktop": false,
                                "is_anon_chat_enabled_for_new_users_only": false,
                                "is_try_it_first_on_login_page_enabled": false,
                                "is_no_auth_welcome_modal_enabled": false,
                                "no_auth_soft_rate_limit": 1200,
                                "no_auth_hard_rate_limit": 1200,
                                "should_show_no_auth_signup_banner": true,
                                "is_no_auth_welcome_back_modal_enabled": true
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ],
                            "explicit_parameters": [
                                "is_anon_chat_enabled"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "101132775": {
                            "name": "101132775",
                            "value": {
                                "show_new_ui": true
                            },
                            "group": "launchedGroup",
                            "rule_id": "launchedGroup",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false
                        },
                        "340466392": {
                            "name": "340466392",
                            "value": {
                                "google.com": [
                                    "google.com",
                                    "googleapis.com",
                                    "www.googleapis.com",
                                    "oauth2.googleapis.com",
                                    "cloudfunctions.net"
                                ],
                                "microsoftonline.com": [
                                    "microsoftonline.com",
                                    "azurewebsites.net",
                                    "graph.microsoft.com",
                                    "microsoft.com",
                                    "azure.com",
                                    "sinequa.com"
                                ],
                                "adobelogin.com": [
                                    "adobe.com",
                                    "adobelogin.com"
                                ],
                                "useflorian.com": [
                                    "useflorian.com",
                                    "convex.site"
                                ],
                                "yahoo.com": [
                                    "yahooapis.com",
                                    "yahoo.com"
                                ],
                                "visualstudio.com": [
                                    "visualstudio.com",
                                    "azure.com"
                                ],
                                "atlassian.com": [
                                    "atlassian.net",
                                    "atlassian.com"
                                ]
                            },
                            "group": "default",
                            "rule_id": "default",
                            "is_device_based": true,
                            "secondary_exposures": []
                        },
                        "445876862": {
                            "name": "445876862",
                            "value": {
                                "is_anon_chat_enabled": true,
                                "anon_composer_display_variant": "default",
                                "anon-is-spanish-translation-enabled": true,
                                "should_show_anon_login_header_on_desktop": false,
                                "is_anon_chat_enabled_for_new_users_only": false,
                                "is_try_it_first_on_login_page_enabled": false,
                                "is_no_auth_welcome_modal_enabled": false,
                                "no_auth_soft_rate_limit": 1200,
                                "no_auth_hard_rate_limit": 1200,
                                "should_show_no_auth_signup_banner": true,
                                "is_no_auth_welcome_back_modal_enabled": true
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ],
                            "explicit_parameters": [
                                "is_anon_chat_enabled",
                                "is_no_auth_welcome_modal_enabled",
                                "no_auth_soft_rate_limit",
                                "no_auth_hard_rate_limit",
                                "should_show_no_auth_signup_banner"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "855925714": {
                            "name": "855925714",
                            "value": {
                                "should_simplify_modal": false,
                                "is_simplified_sharing_modal_enabled": false,
                                "is_social_share_options_enabled": false,
                                "is_update_shared_links_enabled": false,
                                "is_discoverability_toggle_enabled": false,
                                "show_copylink_state_if_no_updates": false
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "explicit_parameters": [
                                "is_simplified_sharing_modal_enabled",
                                "is_social_share_options_enabled",
                                "is_update_shared_links_enabled",
                                "is_discoverability_toggle_enabled"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "889327355": {
                            "name": "889327355",
                            "value": {
                                "is-enabled": true
                            },
                            "group": "layerAssignment",
                            "rule_id": "layerAssignment",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "explicit_parameters": [
                                "is-enabled"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": true,
                            "is_in_layer": true
                        },
                        "1193714591": {
                            "name": "1193714591",
                            "value": {
                                "is_dynamic_model_enabled": false,
                                "show_message_model_info": false,
                                "show_message_regenerate_model_selector": true,
                                "is_conversation_model_switching_allowed": true,
                                "show_rate_limit_downgrade_banner": true,
                                "config": {
                                    "impl": "rm_score",
                                    "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                                    "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                                },
                                "show_message_regenerate_model_selector_on_every_message": true,
                                "is_AG8PqS2q_enabled": false
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "explicit_parameters": [
                                "is_dynamic_model_enabled",
                                "show_message_regenerate_model_selector",
                                "is_conversation_model_switching_allowed",
                                "show_rate_limit_downgrade_banner"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "1193714592": {
                            "name": "1193714592",
                            "value": {
                                "is_dynamic_model_enabled": false,
                                "show_message_model_info": false,
                                "show_message_regenerate_model_selector": true,
                                "is_conversation_model_switching_allowed": true,
                                "show_rate_limit_downgrade_banner": true,
                                "config": {
                                    "impl": "rm_score",
                                    "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                                    "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                                },
                                "show_message_regenerate_model_selector_on_every_message": true,
                                "is_AG8PqS2q_enabled": false
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "explicit_parameters": [
                                "is_dynamic_model_enabled",
                                "show_message_regenerate_model_selector",
                                "is_conversation_model_switching_allowed",
                                "show_rate_limit_downgrade_banner",
                                "config"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "1193714593": {
                            "name": "1193714593",
                            "value": {
                                "is_dynamic_model_enabled": false,
                                "show_message_model_info": false,
                                "show_message_regenerate_model_selector": true,
                                "is_conversation_model_switching_allowed": true,
                                "show_rate_limit_downgrade_banner": true,
                                "config": {
                                    "impl": "rm_score",
                                    "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                                    "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                                },
                                "show_message_regenerate_model_selector_on_every_message": true,
                                "is_AG8PqS2q_enabled": false
                            },
                            "group": "prestart",
                            "rule_id": "prestart",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "explicit_parameters": [
                                "is_dynamic_model_enabled",
                                "show_message_regenerate_model_selector",
                                "is_conversation_model_switching_allowed",
                                "show_rate_limit_downgrade_banner",
                                "config",
                                "show_message_regenerate_model_selector_on_every_message"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "1237904433": {
                            "name": "1237904433",
                            "value": {
                                "is-enabled": true
                            },
                            "group": "launchedGroup",
                            "rule_id": "launchedGroup",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "explicit_parameters": [
                                "is-enabled"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "1277879515": {
                            "name": "1277879515",
                            "value": {
                                "show_ui": true
                            },
                            "group": "launchedGroup",
                            "rule_id": "launchedGroup",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "segment:friends_and_family_nda",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:feather_memory_trainers",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "chatgpt-allow-memory-use-admin",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false
                        },
                        "1312738797": {
                            "name": "1312738797",
                            "value": {
                                "is_anon_chat_enabled": true,
                                "anon_composer_display_variant": "default",
                                "anon-is-spanish-translation-enabled": true,
                                "should_show_anon_login_header_on_desktop": false,
                                "is_anon_chat_enabled_for_new_users_only": false,
                                "is_try_it_first_on_login_page_enabled": false,
                                "is_no_auth_welcome_modal_enabled": false,
                                "no_auth_soft_rate_limit": 1200,
                                "no_auth_hard_rate_limit": 1200,
                                "should_show_no_auth_signup_banner": true,
                                "is_no_auth_welcome_back_modal_enabled": true
                            },
                            "group": "launchedGroup",
                            "rule_id": "launchedGroup",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ],
                            "explicit_parameters": [
                                "anon-is-spanish-translation-enabled"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "1476759596": {
                            "name": "1476759596",
                            "value": {},
                            "group": "layerAssignment",
                            "rule_id": "layerAssignment",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "is_user_in_experiment": false,
                            "is_experiment_active": true
                        },
                        "1547924522": {
                            "name": "1547924522",
                            "value": {},
                            "group": "targetingGate",
                            "rule_id": "targetingGate",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt-paid-users",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false
                        },
                        "1887692490": {
                            "name": "1887692490",
                            "value": {
                                "is_anon_chat_enabled": true,
                                "anon_composer_display_variant": "default",
                                "anon-is-spanish-translation-enabled": true,
                                "should_show_anon_login_header_on_desktop": false,
                                "is_anon_chat_enabled_for_new_users_only": false,
                                "is_try_it_first_on_login_page_enabled": false,
                                "is_no_auth_welcome_modal_enabled": false,
                                "no_auth_soft_rate_limit": 1200,
                                "no_auth_hard_rate_limit": 1200,
                                "should_show_no_auth_signup_banner": true,
                                "is_no_auth_welcome_back_modal_enabled": true
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ],
                            "explicit_parameters": [
                                "is_anon_chat_enabled",
                                "is_try_it_first_on_login_page_enabled"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "1908072088": {
                            "name": "1908072088",
                            "value": {
                                "blocking_modal": true,
                                "single_button_variant": true,
                                "dropdown_tooltip": false,
                                "conversation_model_switcher_badge": false
                            },
                            "group": "default",
                            "rule_id": "default",
                            "is_device_based": false,
                            "secondary_exposures": []
                        },
                        "2144828235": {
                            "name": "2144828235",
                            "value": {},
                            "group": "layerAssignment",
                            "rule_id": "layerAssignment",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "is_user_in_experiment": false,
                            "is_experiment_active": true
                        },
                        "2176179500": {
                            "name": "2176179500",
                            "value": {
                                "should_enable_zh_tw": false
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "explicit_parameters": [
                                "should_enable_zh_tw"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "2331692225": {
                            "name": "2331692225",
                            "value": {},
                            "group": "default",
                            "rule_id": "default",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "chatgpt-browse-sahara",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                }
                            ]
                        },
                        "2370410632": {
                            "name": "2370410632",
                            "value": {
                                "is_anon_chat_enabled": true,
                                "anon_composer_display_variant": "default",
                                "anon-is-spanish-translation-enabled": true,
                                "should_show_anon_login_header_on_desktop": false,
                                "is_anon_chat_enabled_for_new_users_only": false,
                                "is_try_it_first_on_login_page_enabled": false,
                                "is_no_auth_welcome_modal_enabled": false,
                                "no_auth_soft_rate_limit": 1200,
                                "no_auth_hard_rate_limit": 1200,
                                "should_show_no_auth_signup_banner": true,
                                "is_no_auth_welcome_back_modal_enabled": true
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ],
                            "explicit_parameters": [
                                "is_anon_chat_enabled"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "2370410634": {
                            "name": "2370410634",
                            "value": {
                                "is_anon_chat_enabled": true,
                                "anon_composer_display_variant": "default",
                                "anon-is-spanish-translation-enabled": true,
                                "should_show_anon_login_header_on_desktop": false,
                                "is_anon_chat_enabled_for_new_users_only": false,
                                "is_try_it_first_on_login_page_enabled": false,
                                "is_no_auth_welcome_modal_enabled": false,
                                "no_auth_soft_rate_limit": 1200,
                                "no_auth_hard_rate_limit": 1200,
                                "should_show_no_auth_signup_banner": true,
                                "is_no_auth_welcome_back_modal_enabled": true
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ],
                            "explicit_parameters": [
                                "is_anon_chat_enabled"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "2575540972": {
                            "name": "2575540972",
                            "value": {
                                "is_dynamic_model_enabled": false,
                                "show_message_model_info": false,
                                "show_message_regenerate_model_selector": true,
                                "is_conversation_model_switching_allowed": true,
                                "show_rate_limit_downgrade_banner": true,
                                "config": {
                                    "impl": "rm_score",
                                    "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                                    "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                                },
                                "show_message_regenerate_model_selector_on_every_message": true,
                                "is_AG8PqS2q_enabled": false
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "explicit_parameters": [
                                "is_dynamic_model_enabled",
                                "show_message_regenerate_model_selector",
                                "is_conversation_model_switching_allowed",
                                "show_rate_limit_downgrade_banner"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "2576123431": {
                            "name": "2576123431",
                            "value": {
                                "is_anon_chat_enabled": true,
                                "anon_composer_display_variant": "default",
                                "anon-is-spanish-translation-enabled": true,
                                "should_show_anon_login_header_on_desktop": false,
                                "is_anon_chat_enabled_for_new_users_only": false,
                                "is_try_it_first_on_login_page_enabled": false,
                                "is_no_auth_welcome_modal_enabled": false,
                                "no_auth_soft_rate_limit": 1200,
                                "no_auth_hard_rate_limit": 1200,
                                "should_show_no_auth_signup_banner": true,
                                "is_no_auth_welcome_back_modal_enabled": true
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ],
                            "explicit_parameters": [
                                "is_anon_chat_enabled",
                                "is_anon_chat_enabled_for_new_users_only"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "2635734272": {
                            "name": "2635734272",
                            "value": {
                                "should_simplify_modal": false,
                                "is_simplified_sharing_modal_enabled": false,
                                "is_social_share_options_enabled": false,
                                "is_update_shared_links_enabled": false,
                                "is_discoverability_toggle_enabled": false,
                                "show_copylink_state_if_no_updates": false
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "explicit_parameters": [
                                "is_simplified_sharing_modal_enabled",
                                "is_social_share_options_enabled",
                                "is_update_shared_links_enabled",
                                "is_discoverability_toggle_enabled"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "2811617088": {
                            "name": "2811617088",
                            "value": {
                                "is_anon_chat_enabled": true,
                                "anon_composer_display_variant": "default",
                                "anon-is-spanish-translation-enabled": true,
                                "should_show_anon_login_header_on_desktop": false,
                                "is_anon_chat_enabled_for_new_users_only": false,
                                "is_try_it_first_on_login_page_enabled": false,
                                "is_no_auth_welcome_modal_enabled": false,
                                "no_auth_soft_rate_limit": 1200,
                                "no_auth_hard_rate_limit": 1200,
                                "should_show_no_auth_signup_banner": true,
                                "is_no_auth_welcome_back_modal_enabled": true
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ],
                            "explicit_parameters": [
                                "is_anon_chat_enabled",
                                "is_no_auth_welcome_modal_enabled",
                                "no_auth_soft_rate_limit",
                                "no_auth_hard_rate_limit",
                                "should_show_no_auth_signup_banner"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "2919366999": {
                            "name": "2919366999",
                            "value": {
                                "should_simplify_modal": false,
                                "is_simplified_sharing_modal_enabled": false,
                                "is_social_share_options_enabled": false,
                                "is_update_shared_links_enabled": false,
                                "is_discoverability_toggle_enabled": false,
                                "show_copylink_state_if_no_updates": true
                            },
                            "group": "3fdACNmSdaY0SPpdEdJNZI",
                            "rule_id": "3fdACNmSdaY0SPpdEdJNZI",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_web_sharing_modal_simplified_targeting",
                                    "gateValue": "true",
                                    "ruleID": "48mToGJmOtQ5aRx9MCjOor"
                                }
                            ],
                            "explicit_parameters": [
                                "show_copylink_state_if_no_updates"
                            ],
                            "is_user_in_experiment": true,
                            "is_experiment_active": true,
                            "is_in_layer": true
                        },
                        "3168246095": {
                            "name": "3168246095",
                            "value": {
                                "gizmo_ids": []
                            },
                            "group": "default",
                            "rule_id": "default",
                            "is_device_based": false,
                            "secondary_exposures": []
                        },
                        "3299284219": {
                            "name": "3299284219",
                            "value": {
                                "is_dynamic_model_enabled": false,
                                "show_message_model_info": false,
                                "show_message_regenerate_model_selector": true,
                                "is_conversation_model_switching_allowed": true,
                                "show_rate_limit_downgrade_banner": true,
                                "config": {
                                    "impl": "rm_score",
                                    "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                                    "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                                },
                                "show_message_regenerate_model_selector_on_every_message": true,
                                "is_AG8PqS2q_enabled": false
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "explicit_parameters": [
                                "is_dynamic_model_enabled",
                                "show_message_regenerate_model_selector",
                                "is_conversation_model_switching_allowed",
                                "show_rate_limit_downgrade_banner",
                                "show_message_regenerate_model_selector_on_every_message"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "3340924940": {
                            "name": "3340924940",
                            "value": {
                                "is_anon_chat_enabled": true,
                                "anon_composer_display_variant": "default",
                                "anon-is-spanish-translation-enabled": true,
                                "should_show_anon_login_header_on_desktop": false,
                                "is_anon_chat_enabled_for_new_users_only": false,
                                "is_try_it_first_on_login_page_enabled": false,
                                "is_no_auth_welcome_modal_enabled": false,
                                "no_auth_soft_rate_limit": 1200,
                                "no_auth_hard_rate_limit": 1200,
                                "should_show_no_auth_signup_banner": true,
                                "is_no_auth_welcome_back_modal_enabled": true
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ],
                            "explicit_parameters": [
                                "is_anon_chat_enabled",
                                "should_show_anon_login_header_on_desktop"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "3414158273": {
                            "name": "3414158273",
                            "value": {
                                "is_anon_chat_enabled": true,
                                "anon_composer_display_variant": "default",
                                "anon-is-spanish-translation-enabled": true,
                                "should_show_anon_login_header_on_desktop": false,
                                "is_anon_chat_enabled_for_new_users_only": false,
                                "is_try_it_first_on_login_page_enabled": false,
                                "is_no_auth_welcome_modal_enabled": false,
                                "no_auth_soft_rate_limit": 1200,
                                "no_auth_hard_rate_limit": 1200,
                                "should_show_no_auth_signup_banner": true,
                                "is_no_auth_welcome_back_modal_enabled": true
                            },
                            "group": "launchedGroup",
                            "rule_id": "launchedGroup",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ],
                            "explicit_parameters": [
                                "is_try_it_first_on_login_page_enabled",
                                "is_no_auth_welcome_back_modal_enabled"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "3414159263": {
                            "name": "3414159263",
                            "value": {
                                "is_anon_chat_enabled": true,
                                "anon_composer_display_variant": "default",
                                "anon-is-spanish-translation-enabled": true,
                                "should_show_anon_login_header_on_desktop": false,
                                "is_anon_chat_enabled_for_new_users_only": false,
                                "is_try_it_first_on_login_page_enabled": false,
                                "is_no_auth_welcome_modal_enabled": false,
                                "no_auth_soft_rate_limit": 1200,
                                "no_auth_hard_rate_limit": 1200,
                                "should_show_no_auth_signup_banner": true,
                                "is_no_auth_welcome_back_modal_enabled": true
                            },
                            "group": "layerAssignment",
                            "rule_id": "layerAssignment",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ],
                            "explicit_parameters": [
                                "is_no_auth_welcome_back_modal_enabled"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": true,
                            "is_in_layer": true
                        },
                        "3758099832": {
                            "name": "3758099832",
                            "value": {},
                            "group": "prestart",
                            "rule_id": "prestart",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false
                        },
                        "3917128856": {
                            "name": "3917128856",
                            "value": {
                                "is_dynamic_model_enabled": false,
                                "show_message_model_info": false,
                                "show_message_regenerate_model_selector": true,
                                "is_conversation_model_switching_allowed": true,
                                "show_rate_limit_downgrade_banner": true,
                                "config": {
                                    "impl": "rm_score",
                                    "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                                    "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                                },
                                "show_message_regenerate_model_selector_on_every_message": true,
                                "is_AG8PqS2q_enabled": false
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "explicit_parameters": [
                                "is_dynamic_model_enabled",
                                "show_message_regenerate_model_selector",
                                "is_conversation_model_switching_allowed",
                                "show_rate_limit_downgrade_banner",
                                "config"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        },
                        "4204547666": {
                            "name": "4204547666",
                            "value": {
                                "is-enabled": true
                            },
                            "group": "abandoned",
                            "rule_id": "abandoned",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "segment:internal_ips",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "chatgpt-localization-preview",
                                    "gateValue": "false",
                                    "ruleID": "14DZA2LumaPqAdCo52CrUB"
                                },
                                {
                                    "gate": "chatgpt-localization-allowlist",
                                    "gateValue": "false",
                                    "ruleID": "66covmutTYx82FWVUlZAqF"
                                }
                            ],
                            "explicit_parameters": [
                                "is-enabled"
                            ],
                            "is_user_in_experiment": false,
                            "is_experiment_active": false,
                            "is_in_layer": true
                        }
                    },
                    "layer_configs": {
                        "1238742812": {
                            "name": "1238742812",
                            "value": {
                                "should_enable_zh_tw": false
                            },
                            "group": "default",
                            "rule_id": "default",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "explicit_parameters": [],
                            "undelegated_secondary_exposures": []
                        },
                        "1547743984": {
                            "name": "1547743984",
                            "value": {
                                "should_simplify_modal": false,
                                "is_simplified_sharing_modal_enabled": false,
                                "is_social_share_options_enabled": false,
                                "is_update_shared_links_enabled": false,
                                "is_discoverability_toggle_enabled": false,
                                "show_copylink_state_if_no_updates": true
                            },
                            "group": "3fdACNmSdaY0SPpdEdJNZI",
                            "rule_id": "3fdACNmSdaY0SPpdEdJNZI",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_web_sharing_modal_simplified_targeting",
                                    "gateValue": "true",
                                    "ruleID": "48mToGJmOtQ5aRx9MCjOor"
                                }
                            ],
                            "explicit_parameters": [
                                "show_copylink_state_if_no_updates"
                            ],
                            "allocated_experiment_name": "2919366999",
                            "is_experiment_active": true,
                            "is_user_in_experiment": true,
                            "undelegated_secondary_exposures": []
                        },
                        "2723963139": {
                            "name": "2723963139",
                            "value": {
                                "is_dynamic_model_enabled": false,
                                "show_message_model_info": false,
                                "show_message_regenerate_model_selector": true,
                                "is_conversation_model_switching_allowed": true,
                                "show_rate_limit_downgrade_banner": true,
                                "config": {
                                    "impl": "rm_score",
                                    "rm_gpt4_vs_sahara": "rm-model-switcher-1p-d36-0325",
                                    "rm_renderer": "harmony_v4.0.13_8k_turbo_mm"
                                },
                                "show_message_regenerate_model_selector_on_every_message": true,
                                "is_AG8PqS2q_enabled": false
                            },
                            "group": "default",
                            "rule_id": "default",
                            "is_device_based": false,
                            "secondary_exposures": [],
                            "explicit_parameters": [],
                            "undelegated_secondary_exposures": []
                        },
                        "3048336830": {
                            "name": "3048336830",
                            "value": {
                                "is-enabled": true
                            },
                            "group": "default",
                            "rule_id": "default",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "segment:internal_ips",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "chatgpt-localization-preview",
                                    "gateValue": "false",
                                    "ruleID": "14DZA2LumaPqAdCo52CrUB"
                                },
                                {
                                    "gate": "chatgpt-localization-allowlist",
                                    "gateValue": "false",
                                    "ruleID": "66covmutTYx82FWVUlZAqF"
                                }
                            ],
                            "explicit_parameters": [],
                            "undelegated_secondary_exposures": [
                                {
                                    "gate": "segment:internal_ips",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "segment:employees",
                                    "gateValue": "false",
                                    "ruleID": "default"
                                },
                                {
                                    "gate": "chatgpt-localization-preview",
                                    "gateValue": "false",
                                    "ruleID": "14DZA2LumaPqAdCo52CrUB"
                                },
                                {
                                    "gate": "chatgpt-localization-allowlist",
                                    "gateValue": "false",
                                    "ruleID": "66covmutTYx82FWVUlZAqF"
                                }
                            ]
                        },
                        "3637408529": {
                            "name": "3637408529",
                            "value": {
                                "is_anon_chat_enabled": false,
                                "anon_composer_display_variant": "default",
                                "anon-is-spanish-translation-enabled": true,
                                "should_show_anon_login_header_on_desktop": false,
                                "is_anon_chat_enabled_for_new_users_only": false,
                                "is_try_it_first_on_login_page_enabled": false,
                                "is_no_auth_welcome_modal_enabled": false,
                                "no_auth_soft_rate_limit": 1200,
                                "no_auth_hard_rate_limit": 1200,
                                "should_show_no_auth_signup_banner": false,
                                "is_no_auth_welcome_back_modal_enabled": false
                            },
                            "group": "default",
                            "rule_id": "default",
                            "is_device_based": false,
                            "secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ],
                            "explicit_parameters": [],
                            "undelegated_secondary_exposures": [
                                {
                                    "gate": "chatgpt_anon_chat_holdout_20240227",
                                    "gateValue": "false",
                                    "ruleID": "disabled"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_gate_20240328",
                                    "gateValue": "true",
                                    "ruleID": "1f4cSeenpcmWQ9eMKKQhF5"
                                },
                                {
                                    "gate": "chatgpt_no_auth_holdout_20240328",
                                    "gateValue": "false",
                                    "ruleID": "14N13nfiNns0pJGxqO9VVj:2.00:1"
                                }
                            ]
                        }
                    },
                    "sdkParams": {},
                    "has_updates": true,
                    "generator": "statsig-node-sdk",
                    "time": 0,
                    "evaluated_keys": {
                        "userID": "user-xyhelper",
                        "customIDs": {
                            "WebAnonymousCookieID": "9d8b6e32-4f61-4371-8b15-91823f3016cd",
                            "DeviceId": "9d8b6e32-4f61-4371-8b15-91823f3016cd",
                            "stableID": "9d8b6e32-4f61-4371-8b15-91823f3016cd"
                        }
                    },
                    "hash_used": "djb2",
                    "user_hash": "4099313694"
                },
                "user": {
                    "country": "US",
                    "custom": {
                        "auth_status": "logged_in",
                        "has_logged_in_before": true,
                        "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:125.0) Gecko/20100101 Firefox/125.0",
                        "is_paid": false
                    },
                    "customIDs": {
                        "WebAnonymousCookieID": "9d8b6e32-4f61-4371-8b15-91823f3016cd",
                        "DeviceId": "9d8b6e32-4f61-4371-8b15-91823f3016cd",
                        "stableID": "9d8b6e32-4f61-4371-8b15-91823f3016cd"
                    },
                    "ip": "8.8.8.8",
                    "locale": "en",
                    "privateAttributes": {
                        "email": "admin@closeai.biz"
                    },
                    "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:125.0) Gecko/20100101 Firefox/125.0",
                    "userID": "user-xyhelper"
                }
            },
            "authStatus": "logged_in",
            "serverPrimedAllowBrowserStorageValue": true,
            "showCookieConsentBanner": false,
            "isStorageComplianceEnabled": false,
            "ageVerificationDeadline": null,
            "ageVerificationType": "none",
            "logOutUser": false
        },
        "__N_SSP": true
    },
    "page": "/[[...default]]",
    "query": {},
    "buildId": "FUn368sO8Yo1nq60-uOqJ",
    "assetPrefix": "https://cdn.oaistatic.com",
    "isFallback": false,
    "isExperimentalCompile": false,
    "gssp": true,
    "scriptLoader": []
}`
)
