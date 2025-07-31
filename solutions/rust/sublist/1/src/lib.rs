#[derive(Debug, PartialEq, Eq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(_first_list: &[T], _second_list: &[T]) -> Comparison {
    if _first_list == _second_list {
        return Comparison::Equal
    }

    if present(_first_list, _second_list) {
        return Comparison::Sublist
    }
    
    if present(_second_list, _first_list) {
        return Comparison::Superlist
    }

    Comparison::Unequal
}

fn present<T: PartialEq>(first: &[T], second: &[T]) -> bool {
    'outer: for (index, _item) in second.iter().enumerate() {
        for (index2, item2) in first.iter().enumerate() {
            let item = second.get(index + index2);

            match item {
                None => {
                    continue 'outer;
                },
                Some(item) => {
                    if item2 != item {
                        continue 'outer;
                    }
                }
            }
        }

        return true
    }

    false
}