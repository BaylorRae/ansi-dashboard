#!/usr/bin/env python3

import time
from dashboard import Dashboard

dashboard = Dashboard(width=75)
dashboard.render()

time.sleep(1)

dashboard.status = 'preprocessing'
dashboard.item = '138832'
dashboard.step = 'Processing text'
dashboard.completion = 5
dashboard.render()

time.sleep(1)

dashboard.step = 'Creating sequences'
dashboard.completion = 10
dashboard.render()

time.sleep(3)

dashboard.step = 'Extracting entities'
dashboard.completion = 30
dashboard.render()

time.sleep(1)

dashboard.status = 'classifiying'
dashboard.step = 'Running heuristics'
dashboard.completion = 45
dashboard.render()

time.sleep(10)

dashboard.step = 'Running cnn'
dashboard.completion = 56
dashboard.render()

time.sleep(10)

dashboard.step = 'Running lstm'
dashboard.completion = 74
dashboard.render()

time.sleep(10)

dashboard.step = 'Running cnn+lstm'
dashboard.completion = 78
dashboard.render()

time.sleep(3)

dashboard.status = 'post processing'
dashboard.step = 'Merging blank pages'
dashboard.completion = 84
dashboard.render()

time.sleep(1)

dashboard.status = 'post processing'
dashboard.step = 'Applying RNE'
dashboard.completion = 88
dashboard.render()

time.sleep(1)

dashboard.status = 'saving'
dashboard.step = 'Writing output'
dashboard.completion = 95
dashboard.render()

time.sleep(1)

dashboard.status = 'done'
dashboard.step = 'Completed'
dashboard.completion = 100
dashboard.render()

time.sleep(3)

dashboard.status = 'idle'
dashboard.item = ''
dashboard.step = ''
dashboard.completion = 0
dashboard.render()

while True:
    pass
