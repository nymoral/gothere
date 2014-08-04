import sys

variables = {
        "background": "#FFFFFF",
        "dark-fg": "#2C2C2C",
        "focus-border": "#D3643B",
        "outer": "#E9E7D6",
        "user-bg": "#D15A40",
        "dark-text": "#111111",
        "default-name-bg": "#B84A40",
        "game-h": "#FFFCF6",
        "game-nh": "#F9F5E9",
        "table-border": "black",
        "closed-guess-bg" : "#D74820",
        "open-guess-bg" : "#FFF1C9"
        }

def executeTemplate(template):
    for k, v in variables.items():
        field = "[[%s]]" % k
        while template.find(field) != -1:
            template = template.replace(field, v)
    return template

def main():
    print("Building css")
    args = sys.argv
    if len(args) < 2:
        print("Input file not provided")
        return
    if len(args) < 3:
        print("Output file not provided")
        return
    try:
        inF = open(args[1], 'r')
    except FileNotFoundError:
        print("Input file does not exist!")
        return
    outF = open(args[2], 'w')
    
    template = inF.read()

    css = executeTemplate(template)
    
    outF.write(css)
    inF.close()
    outF.close()
    return

if __name__ == "__main__":
    main()

