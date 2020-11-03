Prism.languages.command = {
    'comment': /\/\*[\s\S]*?\*\/|\/\/.*/,
    'string': {
        pattern: /"(?:[^\\"]|\\.)*"|'[^']*'/,
        greedy: true
    },
    'keyword': {
        pattern: /\b-\S*/ig,
        lookbehind: true
    },
    'important': {
        pattern: /(^|[{};\r\n][ \t]*)[a-z_][\w.-]*/i,
        lookbehind: true
    },
    'boolean': /\b(?:false|true)\b/,
    'operator': /\+/,
    'punctuation': /[{};:]/
};

Prism.languages.commandline = Prism.languages.command