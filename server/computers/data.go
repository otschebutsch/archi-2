package computers

import (
	"database/sql"
	//"fmt"
)

type Computer struct {
	Id_c   int  `json:"id_c"`
	Name string `json:"name"`
	CpuCount int `json:"cpuCount"`
	TotalDiskSpace int64 `json:"totalDiskSpace"`
}

type Disk struct {
	Id_d   int  `json:"id"`
	DiskSpace int64 `json:"diskSpace"`
	Id_c int `json:"id_c"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

/*func (s *Store) CreateComputer(name string, cpuCount int) error {
	if len(name) < 0 {
		return fmt.Errorf("computer name is not provided")
	}
	_, err := s.Db.Exec("INSERT INTO computers (name, cpuCount) VALUES ($1, $2)", name)
	return err
}*/

func (s *Store) ListComputers() ([]*Computer, error) {
	rows, err := s.Db.Query("SELECT id_c, name, cpucount, totaldiskspace FROM computers LIMIT 20")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Computer
	for rows.Next() {
		var c Computer
		if err := rows.Scan(&c.Id_c, &c.Name, &c.CpuCount, &c.TotalDiskSpace); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Computer, 0)
	}
	return res, nil
}

func (s *Store) SetDisk() ([]*Computer, error) {
    s.Db.Exec("UPDATE disks SET id_c = '2' WHERE id_d = '1'")
    s.Db.Exec("UPDATE computers SET totaldiskspace=computers.totaldiskspace + disks.diskspace FROM disks WHERE computers.id_c= '2' AND disks.id_d='1'")
		rows, err := s.Db.Query("SELECT name, totaldiskspace FROM computers WHERE id_c = '2'")
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		var res []*Computer
		for rows.Next() {
			var c Computer
			if err := rows.Scan(&c.Name, &c.TotalDiskSpace); err != nil {
				return nil, err
			}
			res = append(res, &c)
		}
		if res == nil {
			res = make([]*Computer, 0)
		}
		return res, nil
}

/*func (s *Store) CreateDisk(DiskSpace int64, id_c int64) error {
	if DiskSpace < 0 {
		return fmt.Errorf("disk space is not provided")
	}
	_, err := s.Db.Exec("INSERT INTO disks (DiskSpace, id_c) VALUES ($1, $2)", DiskSpace)
	return err
}
func (s *Store) ListDisks() ([]*Disk, error) {
	rows, err := s.Db.Query("SELECT id_d, diskSpace, id_c FROM disks LIMIT 200")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []*Disk
	for rows.Next() {
		var d Disk
		if err := rows.Scan(&d.Id_d, &d.DiskSpace, &d.Id_c); err != nil {
			return nil, err
		}
		result = append(result, &d)
	}
	if result == nil {
		result = make([]*Disk, 0)
	}
	return result, nil
}*/
