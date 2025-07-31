// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

pub struct Player {
    pub health: u32,
    pub mana: Option<u32>,
    pub level: u32,
}

impl Player {
    pub fn revive(&self) -> Option<Player> {
        if self.health == 0 {
            return match self.mana {
                Some(_) => {
                    Some(Player {
                        health: 100,
                        mana: Some(100),
                        level: self.level,
                    })
                }
                None => {
                    Some(Player {
                        health: 100,
                        mana: None,
                        level: self.level,
                    })
                }
            }
        }

        None
    }

    pub fn cast_spell(&mut self, mana_cost: u32) -> u32 {
        match self.mana {
            Some(total) => {
                if total < mana_cost {
                    return 0
                }

                self.mana = Some(total - mana_cost);
                mana_cost * 2
            },
            None => {
                if self.health < mana_cost {
                    self.health = 0;
                } else {
                    self.health -= mana_cost;
                }
                
                0
            },
        }
    }
}
