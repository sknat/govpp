// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package classify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RESTHandler(rpc RPCService) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/classify_add_del_session", func(w http.ResponseWriter, req *http.Request) {
		var request = new(ClassifyAddDelSession)
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(b, request); err != nil {
			http.Error(w, "unmarshal data failed", http.StatusBadRequest)
			return
		}
		reply, err := rpc.ClassifyAddDelSession(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	mux.HandleFunc("/classify_add_del_table", func(w http.ResponseWriter, req *http.Request) {
		var request = new(ClassifyAddDelTable)
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(b, request); err != nil {
			http.Error(w, "unmarshal data failed", http.StatusBadRequest)
			return
		}
		reply, err := rpc.ClassifyAddDelTable(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	mux.HandleFunc("/classify_set_interface_ip_table", func(w http.ResponseWriter, req *http.Request) {
		var request = new(ClassifySetInterfaceIPTable)
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(b, request); err != nil {
			http.Error(w, "unmarshal data failed", http.StatusBadRequest)
			return
		}
		reply, err := rpc.ClassifySetInterfaceIPTable(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	mux.HandleFunc("/classify_set_interface_l2_tables", func(w http.ResponseWriter, req *http.Request) {
		var request = new(ClassifySetInterfaceL2Tables)
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(b, request); err != nil {
			http.Error(w, "unmarshal data failed", http.StatusBadRequest)
			return
		}
		reply, err := rpc.ClassifySetInterfaceL2Tables(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	mux.HandleFunc("/classify_table_by_interface", func(w http.ResponseWriter, req *http.Request) {
		var request = new(ClassifyTableByInterface)
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(b, request); err != nil {
			http.Error(w, "unmarshal data failed", http.StatusBadRequest)
			return
		}
		reply, err := rpc.ClassifyTableByInterface(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	mux.HandleFunc("/classify_table_ids", func(w http.ResponseWriter, req *http.Request) {
		var request = new(ClassifyTableIds)
		reply, err := rpc.ClassifyTableIds(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	mux.HandleFunc("/classify_table_info", func(w http.ResponseWriter, req *http.Request) {
		var request = new(ClassifyTableInfo)
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(b, request); err != nil {
			http.Error(w, "unmarshal data failed", http.StatusBadRequest)
			return
		}
		reply, err := rpc.ClassifyTableInfo(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	mux.HandleFunc("/flow_classify_set_interface", func(w http.ResponseWriter, req *http.Request) {
		var request = new(FlowClassifySetInterface)
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(b, request); err != nil {
			http.Error(w, "unmarshal data failed", http.StatusBadRequest)
			return
		}
		reply, err := rpc.FlowClassifySetInterface(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	mux.HandleFunc("/input_acl_set_interface", func(w http.ResponseWriter, req *http.Request) {
		var request = new(InputACLSetInterface)
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(b, request); err != nil {
			http.Error(w, "unmarshal data failed", http.StatusBadRequest)
			return
		}
		reply, err := rpc.InputACLSetInterface(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	mux.HandleFunc("/output_acl_set_interface", func(w http.ResponseWriter, req *http.Request) {
		var request = new(OutputACLSetInterface)
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(b, request); err != nil {
			http.Error(w, "unmarshal data failed", http.StatusBadRequest)
			return
		}
		reply, err := rpc.OutputACLSetInterface(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	mux.HandleFunc("/policer_classify_set_interface", func(w http.ResponseWriter, req *http.Request) {
		var request = new(PolicerClassifySetInterface)
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body failed", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(b, request); err != nil {
			http.Error(w, "unmarshal data failed", http.StatusBadRequest)
			return
		}
		reply, err := rpc.PolicerClassifySetInterface(req.Context(), request)
		if err != nil {
			http.Error(w, "request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rep, err := json.MarshalIndent(reply, "", "  ")
		if err != nil {
			http.Error(w, "marshal failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(rep)
	})
	return http.HandlerFunc(mux.ServeHTTP)
}
