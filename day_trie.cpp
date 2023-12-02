#include <bits/stdc++.h>

#define fastio                        \
    ios_base::sync_with_stdio(false); \
    cin.tie(NULL)

#define endl "\n"

using namespace std;

class Node{
public:
    int value;
    unordered_map<char, Node*> children;

    Node(){
        value = 0;
    }
};

class Trie{
public:
    Node* root;

    int itr = 0;

    void insert(string res, int val){insert(res, 0, root, val);}
    void insert(string res, int idx, Node* head, int val){
        if(idx == res.length()){
            head->value = val;
            return;
        }

        if(head->children[res[idx]] == NULL){
            head->children[res[idx]] = new Node();
        }

        insert(res, idx+1, head->children[res[idx]], val);
    }

    int search(string res, int idx){return search(res, idx, root);}
    int search(string res, int idx, Node* head){

        itr = idx;
        if(head->value){return head->value;}
        if(idx == res.length()){return head->value;}
        if(res[idx] <='9'){return res[idx]-48;}
        if(head->children[res[idx]] == NULL){return 0;}
        
        return search(res, idx+1, head->children[res[idx]]);
    }

    Trie(){
        root = new Node();
    }



};

int main()
{
    fastio;
    
    Trie* obj = new Trie();
    obj->insert("one", 1);
    obj->insert("two", 2);
    obj->insert("three", 3);
    obj->insert("four", 4);
    obj->insert("five", 5);
    obj->insert("six", 6);
    obj->insert("seven", 7);
    obj->insert("eight", 8);
    obj->insert("nine", 9);


    int res = 0;

    ifstream test ("test.txt"); 
    if ( test.is_open() ) { // always check whether the file is open
        while(test.good()){
            string s;
            test >> s; // pipe file's content into stream

            int idx = 0, n = s.length();
            int start = 0, end = 0;

            while(idx<n){
                if(s[idx]<='9'){
                    if(start){
                        end = s[idx]-48;
                    }
                    else{
                        start = s[idx]-48;
                    }
                    idx++;
                } else {
                    int val = obj->search(s, idx);
                    // cout << val << ' ' << idx << "\n";
                    // cout << start << " " << end << "**\n";
                    if(val){
                        if(start){
                            end = val;
                        }
                        else{
                            start = val;
                        }

                        idx++;
                        // idx = (idx == obj->itr ? idx+1 : obj->itr);
                    }
                    else{
                        idx++;   
                    }
                }
            }

            end = (end ? end : start);
            res += 10*start+end;
            cout << 10*start+end << "\n";
        }
    }

    

    cout << res << "--\n";


    return 0;
}