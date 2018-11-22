use timelogStore;

var bulk = db.timelog.initializeUnorderedBulkOp();
bulk.insert({_id: "05", day:"22", pointype :"start", activity:"read"});
bulk.insert({_id: "06", day:"34", pointype :"stop", activity:"study"});
bulk.insert({_id: "07", day:"31", pointype :"start", activity:"listen"});
bulk.insert({_id: "08", day:"44", pointype :"stop", activity:"listen"});

bulk.execute();
