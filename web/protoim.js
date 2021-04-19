var messageBase = null;
var getUserInfoRequest = null;
protobuf.load("protoim.json", function(err, root) {
    if (err) {
        throw err;    
    }
    messageBase = root.lookupType("protoim.MessageBase");
    getUserInfoRequest = root.lookupType("protoim.GetUserInfoRequest");
});