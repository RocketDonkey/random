"""Detect borders on images.

The technique here is based on the answer to:

  http://stackoverflow.com/questions/10985550/detect-if-an-image-has-a-border-
  programmatically-return-boolean

with minor adjustments, mainly just toggling the offset when adjusting for image
inconsistencies.
"""

import glob
import os
import sys

from PIL import Image
from PIL import ImageChops


IMAGE_TYPES = ('*.gif', '*.jpg', '*.jpeg',)


def DetectBorder(image):
  """Determine whether or not an image contains a border.

  Args:
    image: A PIL Image object.

  Returns:
    A boolean representing whether or not the image has a border.
  """
  image_width, image_height = image.size
  bg = Image.new(image.mode, image.size, image.getpixel((0, 0)))
  diff = ImageChops.difference(image, bg)
  # Adjusting the offset to 115 correctly identified one of the clear cases.
  # Could potentially make this a flag in order to test thresholds.
  diff = ImageChops.add(diff, diff, 2.0, -115)
  bbox = diff.getbbox()
  return all((
      # Ensure that the upper-left bounding box coordinate has no value on the
      # X- or Y-axis (non-zero).
      bbox[0], bbox[1],
      # Ensure that the sum of the upper-left X and lower_right X is less than
      # the original width.
      (bbox[0] + bbox[2]) <= image_width,
      # Ensure that the sum of the upper-left Y and lower_right Y is less than
      # the original height.
      (bbox[1] + bbox[3]) <= image_height
  ))


def main():
  """Determine if images have 1px contrasting borders."""
  if len(sys.argv) != 2:
    raise ValueError('Please provide a single directory path to parse')

  directory = sys.argv[1]
  image_files = [file_ for type_ in IMAGE_TYPES
                 for file_ in glob.glob(os.path.join(directory, type_))]
  for index, image_filepath in enumerate(image_files, 1):
    with open(image_filepath, 'rb') as image_file:
      image = Image.open(image_file)
      if image_filepath.endswith('.gif'):
        # Convert .gifs to RGB as a .gif pixel refers to one of the 256 colors
        # in the GIF color palette, and we want to compare in terms of RGB.
        image = image.convert('RGB')
      print 'Image %d: %s %s' % (index, os.path.basename(image_filepath),
                                 DetectBorder(image))

if __name__ == '__main__':
  main()
