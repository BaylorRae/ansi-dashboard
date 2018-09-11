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

