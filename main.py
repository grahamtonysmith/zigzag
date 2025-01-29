import argparse
import os
import sys
import zigzag


def main(data):
    if not os.path.exists(os.path.join(os.getcwd(), 'data', '{}.csv'.format(data))):
        zigzag.say('path to {data} not found')
        sys.exit(1)
    
    zigzag.say('data found')
    sys.exit(0)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--data', type=str, required=True, help='--data=gc')

    args = vars(parser.parse_args())
    
    main(args.get('data'))
