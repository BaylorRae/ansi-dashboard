#!/usr/bin/env python3

import time

class Dashboard:
    status = 'idle'
    item = ''
    step = ''
    completion = 0

    def __init__(self, width):
        self.width = width

    def render(self):
        self.clear_screen()

        self.border()
        self.write_status('Status', self.status)
        self.write_status('Item', self.item)
        self.write_status('Step', self.step)
        self.write_progress()
        self.border()

        print('\r\n', end='')

    def clear_screen(self):
        print('[2J')
        print('[0;0H')

    def border(self):
        self._repeat(
            chars='-',
            prefix='+',
            suffix='+'
        )

    def write_status(self, title, value):
        self._repeat(
            text=f'{title}: {value}',
            prefix='| ',
            suffix='|'
        )

    def write_progress(self):
        prefix = '| '
        suffix = f' {self.completion}% |'
        bars = '#' * int(((self.completion / 100) * self.width) - (len(prefix) + len(suffix)) - 1)
        self._repeat(
            text=bars,
            prefix=prefix,
            suffix=suffix
        )

    def _repeat(self, text=None, chars=' ', prefix='', suffix=''):
        text = '' if text is None else text
        width = self.width - (len(chars) + len(text) + len(prefix) + len(suffix))

        print(prefix, end='')
        print(text, end='')
        print(chars * width, end='')
        print(suffix, end='\r\n')

dashboard = Dashboard(width=75)
dashboard.render()

time.sleep(1)

dashboard.status = 'preprocessing'
dashboard.item = '138832'
dashboard.step = 'Processing text'
dashboard.completion = 5
dashboard.render()

time.sleep(1)

dashboard.status = 'preprocessing'
dashboard.item = '138832'
dashboard.step = 'Creating sequences'
dashboard.completion = 10
dashboard.render()

time.sleep(3)

dashboard.status = 'preprocessing'
dashboard.item = '138832'
dashboard.step = 'Extracting entities'
dashboard.completion = 30
dashboard.render()

time.sleep(1)

dashboard.status = 'classifiying'
dashboard.item = '138832'
dashboard.step = 'Running heuristics'
dashboard.completion = 45
dashboard.render()

time.sleep(10)

dashboard.status = 'classifiying'
dashboard.item = '138832'
dashboard.step = 'Running cnn'
dashboard.completion = 56
dashboard.render()

time.sleep(10)

dashboard.status = 'classifiying'
dashboard.item = '138832'
dashboard.step = 'Running lstm'
dashboard.completion = 74
dashboard.render()

time.sleep(10)

dashboard.status = 'classifiying'
dashboard.item = '138832'
dashboard.step = 'Running cnn+lstm'
dashboard.completion = 78
dashboard.render()

time.sleep(3)

dashboard.status = 'post processing'
dashboard.item = '138832'
dashboard.step = 'Merging blank pages'
dashboard.completion = 84
dashboard.render()

time.sleep(1)

dashboard.status = 'post processing'
dashboard.item = '138832'
dashboard.step = 'Applying RNE'
dashboard.completion = 88
dashboard.render()

time.sleep(1)

dashboard.status = 'saving'
dashboard.item = '138832'
dashboard.step = 'Writing output'
dashboard.completion = 95
dashboard.render()

time.sleep(1)

dashboard.status = 'done'
dashboard.item = '138832'
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
