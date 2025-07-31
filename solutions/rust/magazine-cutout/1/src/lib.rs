// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut map = HashMap::new();

    for word in magazine {
        match map.get(word) {
            Some(count) => map.insert(word, count + 1),
            None => map.insert(word, 1),
        };
    }

    for word in note {
        if !map.contains_key(word) {
            return false
        }

        match map.get(word) {
            Some(1) => map.remove(word),
            Some(count) => map.insert(word, count - 1),
            None => { return false },
        };
    }

    true
}
