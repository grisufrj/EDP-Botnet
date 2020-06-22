from flask import Flask, jsonify,request
app = Flask(__name__)

@app.route('/command',methods=['GET'])
def display_command():
    return '{"args":["ls","-lha"]}'

@app.route('/send',methods=['POST'])
def recv_command():
    data = request.data
    print(data)
    return "200-Ok"

if __name__ == "__main__":
    app.run()
